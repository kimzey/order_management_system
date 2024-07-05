package test

import (
	"context"
	"errors"
	"testing"

	"github.com/kizmey/order_management_system/pkg/interface/aggregation"
	"github.com/kizmey/order_management_system/pkg/interface/entities"
	_productRepository "github.com/kizmey/order_management_system/pkg/repository/product"
	_transactionRepository "github.com/kizmey/order_management_system/pkg/repository/transaction"
	"github.com/kizmey/order_management_system/pkg/service/transaction"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestTransactionService_Create(t *testing.T) {
	mockTransactionRepo := new(_transactionRepository.NewTransactionRepositoryMock)
	mockProductRepo := new(_productRepository.NewProductRepositoryMock)
	service := transaction.NewTransactionServiceImpl(mockTransactionRepo, mockProductRepo)

	mockTransaction := &aggregation.TransactionEcommerce{
		Tranasaction: &entities.Transaction{TransactionID: "1"},
		Product: []entities.Product{
			{ProductID: "1", ProductName: "Product 1", ProductPrice: 500},
			{ProductID: "2", ProductName: "Product 2", ProductPrice: 250},
		},
		AddessProduct: map[string]uint{"1": 1, "2": 2},
	}

	t.Run("successful create transaction", func(t *testing.T) {
		mockProductRepo.On("FindByID", mock.Anything, "1").Return(&mockTransaction.Product[0], nil).Once()
		mockProductRepo.On("FindByID", mock.Anything, "2").Return(&mockTransaction.Product[1], nil).Once()

		mockTransactionRepo.On("Create", mock.Anything, mock.AnythingOfType("*aggregation.TransactionEcommerce")).Return(mockTransaction.Tranasaction, nil).Once()

		ctx := context.Background()
		createdTransaction, err := service.Create(ctx, mockTransaction)

		assert.NoError(t, err)
		assert.NotNil(t, createdTransaction)
		assert.Equal(t, mockTransaction.Tranasaction.TransactionID, createdTransaction.Tranasaction.TransactionID)
		assert.Equal(t, mockTransaction.Tranasaction.SumPrice, createdTransaction.Tranasaction.SumPrice)
		mockTransactionRepo.AssertExpectations(t)
		mockProductRepo.AssertExpectations(t)
	})

	t.Run("create transaction with product error", func(t *testing.T) {
		mockProductRepo.On("FindByID", mock.Anything, "1").Return(&entities.Product{}, errors.New("product error")).Once()

		ctx := context.Background()
		_, err := service.Create(ctx, mockTransaction)

		assert.Error(t, err)
		assert.Equal(t, "product error", err.Error())

		mockProductRepo.AssertExpectations(t)
	})

	t.Run("create transaction with repository error", func(t *testing.T) {
		mockProductRepo.On("FindByID", mock.Anything, "1").Return(&mockTransaction.Product[0], nil).Once()
		mockProductRepo.On("FindByID", mock.Anything, "2").Return(&mockTransaction.Product[1], nil).Once()

		mockTransactionRepo.On("Create", mock.Anything, mock.AnythingOfType("*aggregation.TransactionEcommerce")).Return(&entities.Transaction{}, errors.New("create error")).Once()

		ctx := context.Background()
		_, err := service.Create(ctx, mockTransaction)

		assert.Error(t, err)
		assert.Equal(t, "create error", err.Error())
		mockTransactionRepo.AssertExpectations(t)
		mockProductRepo.AssertExpectations(t)
	})
}

func TestTransactionService_FindAll(t *testing.T) {
	mockTransactionRepo := new(_transactionRepository.NewTransactionRepositoryMock)
	mockProductRepo := new(_productRepository.NewProductRepositoryMock)
	service := transaction.NewTransactionServiceImpl(mockTransactionRepo, mockProductRepo)

	mockTransactions := []entities.Transaction{
		{TransactionID: "1"},
		{TransactionID: "2"},
	}

	t.Run("successful find all transactions", func(t *testing.T) {
		mockTransactionRepo.On("FindAll", mock.Anything).Return(&mockTransactions, nil).Once()

		ctx := context.Background()
		foundTransactions, err := service.FindAll(ctx)

		assert.NoError(t, err)
		assert.NotNil(t, foundTransactions)
		assert.Equal(t, len(mockTransactions), len(*foundTransactions))
		mockTransactionRepo.AssertExpectations(t)
	})

	t.Run("find all transactions with repository error", func(t *testing.T) {
		mockTransactionRepo.On("FindAll", mock.Anything).Return(&[]entities.Transaction{}, errors.New("find all error")).Once()

		ctx := context.Background()
		_, err := service.FindAll(ctx)

		assert.Error(t, err)
		assert.Equal(t, "find all error", err.Error())
		mockTransactionRepo.AssertExpectations(t)
	})
}

func TestTransactionService_FindByID(t *testing.T) {
	mockTransactionRepo := new(_transactionRepository.NewTransactionRepositoryMock)
	mockProductRepo := new(_productRepository.NewProductRepositoryMock)
	service := transaction.NewTransactionServiceImpl(mockTransactionRepo, mockProductRepo)

	mockTransaction := &entities.Transaction{TransactionID: "1"}

	t.Run("successful find transaction by ID", func(t *testing.T) {
		mockTransactionRepo.On("FindByID", mock.Anything, "1").Return(mockTransaction, nil).Once()

		ctx := context.Background()
		foundTransaction, err := service.FindByID(ctx, "1")

		assert.NoError(t, err)
		assert.NotNil(t, foundTransaction)
		assert.Equal(t, mockTransaction.TransactionID, foundTransaction.TransactionID)
		mockTransactionRepo.AssertExpectations(t)
	})

	t.Run("find transaction by ID with repository error", func(t *testing.T) {
		mockTransactionRepo.On("FindByID", mock.Anything, "1").Return(&entities.Transaction{}, errors.New("find by ID error")).Once()

		ctx := context.Background()
		_, err := service.FindByID(ctx, "1")

		assert.Error(t, err)
		assert.Equal(t, "find by ID error", err.Error())
		mockTransactionRepo.AssertExpectations(t)
	})
}
func TestTransactionService_Update(t *testing.T) {
	mockTransactionRepo := new(_transactionRepository.NewTransactionRepositoryMock)
	mockProductRepo := new(_productRepository.NewProductRepositoryMock)
	service := transaction.NewTransactionServiceImpl(mockTransactionRepo, mockProductRepo)

	mockTransaction := &aggregation.TransactionEcommerce{
		Tranasaction: &entities.Transaction{TransactionID: "1"},
		Product: []entities.Product{
			{ProductID: "1", ProductName: "Product 1", ProductPrice: 500},
		},
		AddessProduct: map[string]uint{"1": 1},
	}

	t.Run("successful update transaction", func(t *testing.T) {
		mockProductRepo.On("FindByID", mock.Anything, "1").Return(&mockTransaction.Product[0], nil).Once()

		mockTransactionRepo.On("Update", mock.Anything, "1", mockTransaction).Return(mockTransaction.Tranasaction, nil).Once()

		ctx := context.Background()
		updatedTransaction, err := service.Update(ctx, "1", mockTransaction)

		assert.NoError(t, err)
		assert.NotNil(t, updatedTransaction)
		assert.Equal(t, mockTransaction.Tranasaction.TransactionID, updatedTransaction.Tranasaction.TransactionID)
		mockTransactionRepo.AssertExpectations(t)
		mockProductRepo.AssertExpectations(t)
	})

	t.Run("update transaction with product error", func(t *testing.T) {
		mockProductRepo.On("FindByID", mock.Anything, "1").Return(nil, errors.New("product error")).Once()

		ctx := context.Background()
		_, err := service.Update(ctx, "1", mockTransaction)

		assert.Error(t, err)
		assert.Equal(t, "product error", err.Error())
		mockTransactionRepo.AssertExpectations(t)
		mockProductRepo.AssertExpectations(t)
	})

	t.Run("update transaction with repository error", func(t *testing.T) {
		mockProductRepo.On("FindByID", mock.Anything, "1").Return(&mockTransaction.Product[0], nil).Once()

		mockTransactionRepo.On("Update", mock.Anything, "1", mockTransaction).Return(&entities.Transaction{}, errors.New("update error")).Once()

		ctx := context.Background()
		_, err := service.Update(ctx, "1", mockTransaction)

		assert.Error(t, err)
		assert.Equal(t, "update error", err.Error())
		mockTransactionRepo.AssertExpectations(t)
		mockProductRepo.AssertExpectations(t)
	})
}

func TestTransactionService_Delete(t *testing.T) {
	mockTransactionRepo := new(_transactionRepository.NewTransactionRepositoryMock)
	mockProductRepo := new(_productRepository.NewProductRepositoryMock)
	service := transaction.NewTransactionServiceImpl(mockTransactionRepo, mockProductRepo)

	mockTransaction := &entities.Transaction{TransactionID: "1"}

	t.Run("successful delete transaction", func(t *testing.T) {
		mockTransactionRepo.On("Delete", mock.Anything, "1").Return(mockTransaction, nil).Once()

		ctx := context.Background()
		deletedTransaction, err := service.Delete(ctx, "1")

		assert.NoError(t, err)
		assert.NotNil(t, deletedTransaction)
		assert.Equal(t, mockTransaction.TransactionID, deletedTransaction.TransactionID)
		mockTransactionRepo.AssertExpectations(t)
	})

	t.Run("delete transaction with repository error", func(t *testing.T) {
		mockTransactionRepo.On("Delete", mock.Anything, "1").Return(&entities.Transaction{}, errors.New("delete error")).Once()

		ctx := context.Background()
		_, err := service.Delete(ctx, "1")

		assert.Error(t, err)
		assert.Equal(t, "delete error", err.Error())
		mockTransactionRepo.AssertExpectations(t)
	})
}
