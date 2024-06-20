package service

import (
	"context"
	_interface "github.com/kizmey/order_management_system/pkg/interface"
	"github.com/kizmey/order_management_system/pkg/interface/entities"
	"github.com/kizmey/order_management_system/pkg/interface/modelReq"
	"github.com/kizmey/order_management_system/pkg/interface/modelRes"
	_TransactionRepository "github.com/kizmey/order_management_system/pkg/repository"
)

type transactionService struct {
	transactionRepository _TransactionRepository.TransactionRepository
	stockRepository       _TransactionRepository.StockRepository
	productRepository     _TransactionRepository.ProductRepository
}

func NewTransactionServiceImpl(
	transactionRepository _TransactionRepository.TransactionRepository,
	productRepository _TransactionRepository.ProductRepository,
) TransactionService {
	return &transactionService{
		transactionRepository: transactionRepository,
		productRepository:     productRepository,
	}
}

func (s *transactionService) Create(ctx context.Context, transaction *modelReq.Transaction) (*modelRes.Transaction, error) {

	transactionEntity := s.transactionReqToEntity(transaction)

	var products []entities.Product

	for productID, quantity := range transaction.Product {
		product, err := s.productRepository.FindByID(productID)
		if err != nil {
			return nil, err
		}
		products = append(products, *product)
		transactionEntity.SumPrice += transactionEntity.CalculatePrice(product.ProductPrice, quantity, transactionEntity.IsDomestic)
	}

	transactionEcommerce := _interface.NewTransactionEcommerce(transactionEntity, products, transaction.Product)
	transactionEntity, err := s.transactionRepository.Create(transactionEcommerce)
	if err != nil {
		return nil, err
	}

	transactionRes := s.transactionEntityToRes(transactionEntity)
	findproducs, err := s.transactionRepository.FindProductsByTransactionID(transactionEntity.TransactionID)
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

	transactionEntities, err := s.transactionRepository.FindAll()
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
	transactionEntity, err := s.transactionRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return s.transactionEntityToRes(transactionEntity), nil
}

func (s *transactionService) Update(ctx context.Context, id string, transaction *modelReq.Transaction) (*modelRes.Transaction, error) {
	transactionEntity := s.transactionReqToEntity(transaction)

	var products []entities.Product

	for productID, quantity := range transaction.Product {
		product, err := s.productRepository.FindByID(productID)
		if err != nil {
			return nil, err
		}
		products = append(products, *product)
		transactionEntity.SumPrice += transactionEntity.CalculatePrice(product.ProductPrice, quantity, transactionEntity.IsDomestic)
	}

	transactionEcommerce := _interface.NewTransactionEcommerce(transactionEntity, products, transaction.Product)
	transactionEntity, err := s.transactionRepository.Update(id, transactionEcommerce)
	if err != nil {
		return nil, err
	}

	transactionRes := s.transactionEntityToRes(transactionEntity)
	findproducs, err := s.transactionRepository.FindProductsByTransactionID(transactionEntity.TransactionID)
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

	transaction, err := s.transactionRepository.Delete(id)
	if err != nil {
		return nil, err
	}

	return s.transactionEntityToRes(transaction), nil
}

func (s *transactionService) transactionReqToEntity(transactionReq *modelReq.Transaction) *entities.Transaction {
	productid := make([]string, 0, len(transactionReq.Product))
	quantity := make([]uint, 0, len(transactionReq.Product))

	for key, value := range transactionReq.Product {
		productid = append(productid, key)
		quantity = append(quantity, value)
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
