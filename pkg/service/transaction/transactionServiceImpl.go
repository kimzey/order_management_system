package transaction

import (
	"context"
	"github.com/kizmey/order_management_system/pkg/interface/aggregation"
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
