package transaction

import (
	"context"
	_interface "github.com/kizmey/order_management_system/pkg/interface"
	"github.com/kizmey/order_management_system/pkg/interface/entities"
	"github.com/kizmey/order_management_system/pkg/interface/modelReq"
	"github.com/kizmey/order_management_system/pkg/interface/modelRes"
	_productRepository "github.com/kizmey/order_management_system/pkg/repository/product"
	_transactionRepository "github.com/kizmey/order_management_system/pkg/repository/transaction"
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

func (s *transactionService) Create(ctx context.Context, transaction *modelReq.Transaction) (*modelRes.Transaction, error) {
	ctx, sp := tracer.Start(ctx, "transactionCreateService")
	defer sp.End()

	productMap := make(map[string]uint)
	for _, item := range transaction.Product {
		productMap[item.ProductID] = item.Quantity
	}

	transactionEntity := s.transactionReqToEntity(transaction)

	var products []entities.Product

	for productID, quantity := range productMap {

		product, err := s.productRepository.FindByID(ctx, productID)
		if err != nil {
			return nil, err
		}
		products = append(products, *product)
		transactionEntity.SumPrice += transactionEntity.CalculatePrice(product.ProductPrice, quantity, transactionEntity.IsDomestic)
	}

	transactionEcommerce := _interface.NewTransactionEcommerce(transactionEntity, products, productMap)
	transactionEntity, err := s.transactionRepository.Create(ctx, transactionEcommerce)
	if err != nil {
		return nil, err
	}

	transactionRes := s.transactionEntityToRes(transactionEntity)
	findproducs, err := s.transactionRepository.FindProductsByTransactionID(ctx, transactionEntity.TransactionID)
	if err != nil {
		return nil, err
	}
	for i, product := range findproducs.Product {
		transactionRes.Products = append(transactionRes.Products, modelRes.Product{
			ProductID:    product.ProductID,
			ProductName:  product.ProductName,
			ProductPrice: product.ProductPrice,
			Quantity:     (findproducs.Quantity)[i],
		})
	}

	return transactionRes, nil

}

func (s *transactionService) FindAll(ctx context.Context) (*[]modelRes.Transaction, error) {
	ctx, sp := tracer.Start(ctx, "transactionFindAllService")
	defer sp.End()

	transactionEntities, err := s.transactionRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	allTransaction := make([]modelRes.Transaction, 0)
	for _, transactionEntity := range *transactionEntities {
		allTransaction = append(allTransaction, *s.transactionEntityToRes(&transactionEntity))
	}

	return &allTransaction, nil
}

func (s *transactionService) FindByID(ctx context.Context, id string) (*modelRes.Transaction, error) {
	ctx, sp := tracer.Start(ctx, "transactionFindByIdService")
	defer sp.End()

	transactionEntity, err := s.transactionRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return s.transactionEntityToRes(transactionEntity), nil
}

func (s *transactionService) Update(ctx context.Context, id string, transaction *modelReq.Transaction) (*modelRes.Transaction, error) {
	ctx, sp := tracer.Start(ctx, "transactionUpdateService")
	defer sp.End()

	productMap := make(map[string]uint)
	for _, item := range transaction.Product {
		productMap[item.ProductID] = item.Quantity
	}

	transactionEntity := s.transactionReqToEntity(transaction)

	var products []entities.Product

	for productID, quantity := range productMap {
		product, err := s.productRepository.FindByID(ctx, productID)
		if err != nil {
			return nil, err
		}
		products = append(products, *product)
		transactionEntity.SumPrice += transactionEntity.CalculatePrice(product.ProductPrice, quantity, transactionEntity.IsDomestic)
	}

	transactionEcommerce := _interface.NewTransactionEcommerce(transactionEntity, products, productMap)
	transactionEntity, err := s.transactionRepository.Update(ctx, id, transactionEcommerce)
	if err != nil {
		return nil, err
	}

	transactionRes := s.transactionEntityToRes(transactionEntity)
	findproducs, err := s.transactionRepository.FindProductsByTransactionID(ctx, transactionEntity.TransactionID)
	if err != nil {
		return nil, err
	}
	for i, product := range findproducs.Product {
		transactionRes.Products = append(transactionRes.Products, modelRes.Product{
			ProductID:    product.ProductID,
			ProductName:  product.ProductName,
			ProductPrice: product.ProductPrice,
			Quantity:     (findproducs.Quantity)[i],
		})
	}

	return transactionRes, nil
}

func (s *transactionService) Delete(ctx context.Context, id string) (*modelRes.Transaction, error) {
	ctx, sp := tracer.Start(ctx, "transactionDeleteService")
	defer sp.End()

	transaction, err := s.transactionRepository.Delete(ctx, id)
	if err != nil {
		return nil, err
	}

	return s.transactionEntityToRes(transaction), nil
}

func (s *transactionService) transactionReqToEntity(transactionReq *modelReq.Transaction) *entities.Transaction {

	productid := make([]string, 0, len(transactionReq.Product))
	quantity := make([]uint, 0, len(transactionReq.Product))

	for _, item := range transactionReq.Product {
		productid = append(productid, item.ProductID)
		quantity = append(quantity, item.Quantity)
	}

	entityProduct := &entities.Transaction{
		IsDomestic: transactionReq.IsDomestic,
	}
	return entityProduct
}

func (s *transactionService) transactionEntityToRes(transactionEntity *entities.Transaction) *modelRes.Transaction {
	return &modelRes.Transaction{
		TransactionID: transactionEntity.TransactionID,
		IsDomestic:    transactionEntity.IsDomestic,
		SumPrice:      transactionEntity.SumPrice,
	}
}
