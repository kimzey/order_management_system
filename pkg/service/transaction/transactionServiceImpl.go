package transaction

import (
	"context"
	_interface "github.com/kizmey/order_management_system/pkg/interface"
	"github.com/kizmey/order_management_system/pkg/interface/entities"
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

func (s *transactionService) Create(ctx context.Context, transaction *_interface.TransactionEcommerce,
) (*_interface.TransactionEcommerce, error) {
	ctx, sp := tracer.Start(ctx, "transactionCreateService")
	defer sp.End()

	var products []entities.Product

	for productID, quantity := range transaction.AddessProduct {
		product, err := s.productRepository.FindByID(ctx, productID)
		if err != nil {
			return nil, err
		}
		products = append(products, *product)
		transaction.Tranasaction.SumPrice += transaction.Tranasaction.CalculatePrice(product.ProductPrice, quantity, transaction.Tranasaction.IsDomestic)
	}

	transactionEcommerce := _interface.NewTransactionEcommerce(transaction.Tranasaction, products, transaction.AddessProduct)
	transactionEntity, err := s.transactionRepository.Create(ctx, transactionEcommerce)
	if err != nil {
		return nil, err
	}
	transaction.Tranasaction = transactionEntity
	findproducs, err := s.transactionRepository.FindProductsByTransactionID(ctx, transactionEntity.TransactionID)
	if err != nil {
		return nil, err
	}
	for _, product := range findproducs.Product {
		transaction.Product = append(transaction.Product, product)
	}

	return transaction, nil

}

func (s *transactionService) FindAll(ctx context.Context) (*[]entities.Transaction, error) {
	ctx, sp := tracer.Start(ctx, "transactionFindAllService")
	defer sp.End()

	transactionEntities, err := s.transactionRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return transactionEntities, nil
}

func (s *transactionService) FindByID(ctx context.Context, id string) (*entities.Transaction, error) {
	ctx, sp := tracer.Start(ctx, "transactionFindByIdService")
	defer sp.End()

	transactionEntity, err := s.transactionRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return transactionEntity, nil
}

func (s *transactionService) Update(ctx context.Context, id string, transaction *_interface.TransactionEcommerce,
) (*_interface.TransactionEcommerce, error) {
	ctx, sp := tracer.Start(ctx, "transactionUpdateService")
	defer sp.End()

	var products []entities.Product

	for productID, quantity := range transaction.AddessProduct {
		product, err := s.productRepository.FindByID(ctx, productID)
		if err != nil {
			return nil, err
		}
		products = append(products, *product)
		transaction.Tranasaction.SumPrice += transaction.Tranasaction.CalculatePrice(product.ProductPrice, quantity, transaction.Tranasaction.IsDomestic)
	}

	transactionEcommerce := _interface.NewTransactionEcommerce(transaction.Tranasaction, products, transaction.AddessProduct)
	transactionEntity, err := s.transactionRepository.Update(ctx, id, transactionEcommerce)
	if err != nil {
		return nil, err
	}

	transaction.Tranasaction = transactionEntity
	findproducs, err := s.transactionRepository.FindProductsByTransactionID(ctx, transaction.Tranasaction.TransactionID)
	if err != nil {
		return nil, err
	}
	for _, product := range findproducs.Product {
		transaction.Product = append(transaction.Product, product)
	}

	return transaction, nil
}

func (s *transactionService) Delete(ctx context.Context, id string) (*entities.Transaction, error) {
	ctx, sp := tracer.Start(ctx, "transactionDeleteService")
	defer sp.End()

	transaction, err := s.transactionRepository.Delete(ctx, id)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}
