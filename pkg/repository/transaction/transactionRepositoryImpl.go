package transaction

import (
	"context"
	"errors"
	"fmt"
	"github.com/kizmey/order_management_system/database"
	"github.com/kizmey/order_management_system/pkg/interface/aggregation"
	"github.com/kizmey/order_management_system/pkg/interface/entities"
	"github.com/kizmey/order_management_system/pkg/interface/model"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type transactionRepositoryImpl struct {
	db database.Database
}

func NewTransactionRepositoryImpl(db database.Database) TransactionRepository {
	return &transactionRepositoryImpl{db: db}
}

func (r *transactionRepositoryImpl) Create(ctx context.Context, transaction *aggregation.TransactionEcommerce) (*entities.Transaction, error) {
	_, sp := tracer.Start(ctx, "transactionCreateRepository")
	defer sp.End()

	transactionModel := r.ToTransactionModel(transaction)

	if err := r.db.Connect().Create(&transactionModel).Error; err != nil {
		return nil, errors.New(fmt.Sprintf("failed to create transaction"))
	}

	for productID, quantity := range transaction.AddessProduct {
		if err := r.db.Connect().Model(&model.TransactionProduct{}).Where("transaction_id = ? AND product_id = ?", transactionModel.ID, productID).Update("quantity", quantity).Error; err != nil {
			return nil, errors.New(fmt.Sprintf("failed to update transaction"))
		}
	}
	transactionEntity := transactionModel.ToTransactionEntity()
	r.SetTranactionSubAttributes(transactionEntity, sp)
	return transactionEntity, nil
}

func (r *transactionRepositoryImpl) FindAll(ctx context.Context) (*[]entities.Transaction, error) {
	_, sp := tracer.Start(ctx, "transactionFindAllRepository")
	defer sp.End()

	transactions := new([]model.Transaction)

	if err := r.db.Connect().Find(&transactions).Error; err != nil {
		return nil, errors.New(fmt.Sprintf("failed to find transactions"))
	}
	allTransactions := model.ConvertModelsTransactionToEntities(transactions)

	r.SetTranactionSubAttributes(allTransactions, sp)
	return allTransactions, nil
}

func (r *transactionRepositoryImpl) FindByID(ctx context.Context, id string) (*entities.Transaction, error) {
	_, sp := tracer.Start(ctx, "transactionFindByIdRepository")
	defer sp.End()

	transaction := new(model.Transaction)
	if err := r.db.Connect().Where("id = ?", id).First(&transaction).Error; err != nil {
		return nil, errors.New(fmt.Sprintf("failed to find transaction"))
	}

	transactionEntity := transaction.ToTransactionEntity()
	r.SetTranactionSubAttributes(transactionEntity, sp)
	return transactionEntity, nil
}

func (r *transactionRepositoryImpl) Update(ctx context.Context, id string, transaction *aggregation.TransactionEcommerce) (*entities.Transaction, error) {
	_, sp := tracer.Start(ctx, "transactionUpdateRepository")
	defer sp.End()

	transactionModel := r.ToTransactionModel(transaction)

	transactionModel.ID = id
	fmt.Println(transactionModel)
	if err := r.db.Connect().Model(&model.Transaction{}).Where("id = ?", id).Updates(&transactionModel).Error; err != nil {
		return nil, errors.New(fmt.Sprintf("failed to update transaction"))
	}

	for productID, quantity := range transaction.AddessProduct {
		if err := r.db.Connect().Model(&model.TransactionProduct{}).
			Where("transaction_id = ? AND product_id = ?", id, productID).
			Update("quantity", quantity).Error; err != nil {
			return nil, errors.New(fmt.Sprintf("failed to update transaction"))
		}
	}

	transactionEntity := transactionModel.ToTransactionEntity()
	r.SetTranactionSubAttributes(transactionEntity, sp)
	return transactionEntity, nil
}

func (r *transactionRepositoryImpl) Delete(ctx context.Context, id string) (*entities.Transaction, error) {
	_, sp := tracer.Start(ctx, "transactionDeleteRepository")
	defer sp.End()

	transaction := new(model.Transaction)
	if err := r.db.Connect().Where("id = ?", id).First(&transaction).Delete(&transaction).Error; err != nil {
		return nil, errors.New(fmt.Sprintf("failed to delete transaction"))
	}

	transactionEntity := transaction.ToTransactionEntity()
	r.SetTranactionSubAttributes(transactionEntity, sp)
	return transactionEntity, nil
}

func (r *transactionRepositoryImpl) FindProductsByTransactionID(ctx context.Context, id string) (*aggregation.Ecommerce, error) {
	_, sp := tracer.Start(ctx, "transactionFindProductsByTransactionIDRepository")
	defer sp.End()

	var transactionProducts []model.TransactionProduct
	if err := r.db.Connect().Where("transaction_id = ?", id).Preload("Product").Find(&transactionProducts).Error; err != nil {
		return nil, errors.New(fmt.Sprintf("failed to find transaction"))
	}

	var products []entities.Product
	var quantity []uint
	for _, transactionProduct := range transactionProducts {
		products = append(products, *transactionProduct.Product.ToProductEntity())
		quantity = append(quantity, transactionProduct.Quantity)
	}

	ecommerceProducts := aggregation.NewEcommerce(nil, products, quantity)
	r.SetEcommerceSubAttributes(ecommerceProducts, sp)
	return ecommerceProducts, nil
}

func (r *transactionRepositoryImpl) ToTransactionModel(e *aggregation.TransactionEcommerce) *model.Transaction {
	var productlist []model.Product
	for _, v := range e.Product {
		productlist = append(productlist, model.Product{
			ID:    v.ProductID,
			Name:  v.ProductName,
			Price: v.ProductPrice,
		})
	}
	return &model.Transaction{
		SumPrice:   e.Tranasaction.SumPrice,
		IsDomestic: e.Tranasaction.IsDomestic,
		Products:   productlist,
	}
}

func (r *transactionRepositoryImpl) SetTranactionSubAttributes(tranasactionData any, sp trace.Span) {
	if transactions, ok := tranasactionData.(*[]entities.Transaction); ok {
		TransactionIDs := make([]string, len(*transactions))
		SumPrices := make([]int, len(*transactions))
		IsDometic := make([]bool, len(*transactions))

		for _, transaction := range *transactions {
			TransactionIDs = append(TransactionIDs, transaction.TransactionID)
			SumPrices = append(SumPrices, int(transaction.SumPrice))
			IsDometic = append(IsDometic, transaction.IsDomestic)
		}

		sp.SetAttributes(
			attribute.StringSlice("TransactionID", TransactionIDs),
			attribute.IntSlice("SumPrice", SumPrices),
			attribute.BoolSlice("IsDomestic", IsDometic),
		)

	} else if transaction, ok := tranasactionData.(*entities.Transaction); ok {
		sp.SetAttributes(
			attribute.String("TransactionID", transaction.TransactionID),
			attribute.Int("SumPrice", int(transaction.SumPrice)),
			attribute.Bool("IsDomestic", transaction.IsDomestic),
		)
	} else {
		sp.RecordError(errors.New("invalid type"))
	}
}

func (r *transactionRepositoryImpl) SetEcommerceSubAttributes(ecommerceData any, sp trace.Span) {
	if ecommerce, ok := ecommerceData.(*aggregation.Ecommerce); ok {
		quantity := make([]int, len(ecommerce.Quantity))
		productIDs := make([]string, len(ecommerce.Product))
		productNames := make([]string, len(ecommerce.Product))
		productPrices := make([]int, len(ecommerce.Product))

		for _, product := range ecommerce.Product {
			productIDs = append(productIDs, product.ProductID)
			productNames = append(productNames, product.ProductName)
			productPrices = append(productPrices, int(product.ProductPrice))
		}

		for _, data := range ecommerce.Quantity {
			quantity = append(quantity, int(data))
		}

		sp.SetAttributes(
			attribute.IntSlice("Quantity", quantity),
			attribute.StringSlice("ProductIDs", productIDs),
			attribute.StringSlice("ProductNames", productNames),
			attribute.IntSlice("ProductPrices", productPrices),
		)

	} else {
		sp.RecordError(errors.New("invalid type"))
	}
}
