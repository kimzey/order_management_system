package transaction

import (
	"context"
	"errors"
	"fmt"
	"github.com/kizmey/order_management_system/pkg/interface/aggregation"
	"github.com/kizmey/order_management_system/pkg/interface/entities"
	_productRepository "github.com/kizmey/order_management_system/pkg/repository/product"
	_transactionRepository "github.com/kizmey/order_management_system/pkg/repository/transaction"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"strconv"
)

type transactionService struct {
	transactionRepository _transactionRepository.TransactionRepository
	productRepository     _productRepository.ProductRepository
}

func NewTransactionServiceImpl(
	transactionRepository _transactionRepository.TransactionRepository,
	productRepository _productRepository.ProductRepository,
) TransactionService {
	return &transactionService{
		transactionRepository: transactionRepository,
		productRepository:     productRepository,
	}
}

func (s *transactionService) Create(ctx context.Context, transaction *aggregation.TransactionEcommerce,
) (*aggregation.TransactionEcommerce, error) {

	ctx, sp := tracer.Start(ctx, "transactionCreateService")
	defer sp.End()

	for productID := range transaction.AddessProduct {
		product, err := s.productRepository.FindByID(ctx, productID)
		if err != nil {
			return nil, err
		}
		transaction.Product = append(transaction.Product, *product)
	}

	transaction.CalculatePrice()

	transactionEntity, err := s.transactionRepository.Create(ctx, transaction)
	if err != nil {
		return nil, err
	}

	transaction.Tranasaction = transactionEntity
	s.SetTransactionEcommerceSubAttributes(transaction, sp)
	return transaction, nil
}

func (s *transactionService) FindAll(ctx context.Context) (*[]entities.Transaction, error) {
	ctx, sp := tracer.Start(ctx, "transactionFindAllService")
	defer sp.End()

	transactionEntities, err := s.transactionRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	s.SetTranactionSubAttributes(transactionEntities, sp)
	return transactionEntities, nil
}

func (s *transactionService) FindByID(ctx context.Context, id string) (*entities.Transaction, error) {
	ctx, sp := tracer.Start(ctx, "transactionFindByIdService")
	defer sp.End()

	transactionEntity, err := s.transactionRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	s.SetTranactionSubAttributes(transactionEntity, sp)
	return transactionEntity, nil
}

func (s *transactionService) Update(ctx context.Context, id string, transaction *aggregation.TransactionEcommerce,
) (*aggregation.TransactionEcommerce, error) {

	ctx, sp := tracer.Start(ctx, "transactionUpdateService")
	defer sp.End()

	for productID := range transaction.AddessProduct {
		product, err := s.productRepository.FindByID(ctx, productID)
		if err != nil {
			return nil, err
		}
		transaction.Product = append(transaction.Product, *product)
	}

	transaction.CalculatePrice()

	transactionEntity, err := s.transactionRepository.Update(ctx, id, transaction)
	if err != nil {
		return nil, err
	}

	transaction.Tranasaction = transactionEntity
	s.SetTransactionEcommerceSubAttributes(transaction, sp)
	return transaction, nil
}

func (s *transactionService) Delete(ctx context.Context, id string) (*entities.Transaction, error) {
	ctx, sp := tracer.Start(ctx, "transactionDeleteService")
	defer sp.End()

	transaction, err := s.transactionRepository.Delete(ctx, id)
	if err != nil {
		return nil, err
	}

	s.SetTranactionSubAttributes(transaction, sp)
	return transaction, nil
}

func (s *transactionService) SetTranactionSubAttributes(tranasctionData any, sp trace.Span) {
	if transactions, ok := tranasctionData.(*[]entities.Transaction); ok {
		var TransactionIDs []string
		var SumPrices []int
		var IsDometic []bool

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

	} else if transaction, ok := tranasctionData.(*entities.Transaction); ok {
		sp.SetAttributes(
			attribute.String("TransactionID", transaction.TransactionID),
			attribute.Int("SumPrice", int(transaction.SumPrice)),
			attribute.Bool("IsDomestic", transaction.IsDomestic),
		)
	} else {
		sp.RecordError(errors.New("invalid type"))
	}
}

func (s *transactionService) SetTransactionEcommerceSubAttributes(TransactionEcommerceData any, sp trace.Span) {
	if transaction, ok := TransactionEcommerceData.(*aggregation.TransactionEcommerce); ok {

		addressProducts := make([]string, 0, len(transaction.AddessProduct))
		productIds := make([]string, 0, len(transaction.Product))
		productNames := make([]string, 0, len(transaction.Product))
		productPrices := make([]int, 0, len(transaction.Product))

		fmt.Println(productPrices)

		for _, product := range transaction.Product {
			productIds = append(productIds, product.ProductID)
			productNames = append(productNames, product.ProductName)
			productPrices = append(productPrices, int(product.ProductPrice))
		}

		for key, value := range transaction.AddessProduct {
			addressProducts = append(addressProducts, key+" : "+strconv.Itoa(int(value)))
		}

		sp.SetAttributes(
			attribute.String("TransactionID", transaction.Tranasaction.TransactionID),
			attribute.Int("SumPrice", int(transaction.Tranasaction.SumPrice)),
			attribute.Bool("IsDomestic", transaction.Tranasaction.IsDomestic),
			attribute.StringSlice("AddressProducts", addressProducts),
			attribute.StringSlice("ProductIds", productIds),
			attribute.StringSlice("ProductNames", productNames),
			attribute.IntSlice("PrductPrices", productPrices),
		)
	} else {
		sp.RecordError(errors.New("invalid type"))
	}
}
