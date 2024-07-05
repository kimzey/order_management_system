package transaction

import (
	"context"
	"github.com/kizmey/order_management_system/pkg/interface/aggregation"
	"github.com/kizmey/order_management_system/pkg/interface/entities"
	"github.com/stretchr/testify/mock"
)

type NewTransactionRepositoryMock struct {
	mock.Mock
}

func (m *NewTransactionRepositoryMock) transactionRepositoryMock() TransactionRepository {
	return m
}

func (m *NewTransactionRepositoryMock) Create(ctx context.Context, transaction *aggregation.TransactionEcommerce) (*entities.Transaction, error) {
	args := m.Called(ctx, transaction)
	return args.Get(0).(*entities.Transaction), args.Error(1)
}

func (m *NewTransactionRepositoryMock) FindAll(ctx context.Context) (*[]entities.Transaction, error) {
	args := m.Called(ctx)
	return args.Get(0).(*[]entities.Transaction), args.Error(1)
}

func (m *NewTransactionRepositoryMock) FindByID(ctx context.Context, id string) (*entities.Transaction, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*entities.Transaction), args.Error(1)
}

func (m *NewTransactionRepositoryMock) Update(ctx context.Context, id string, transaction *aggregation.TransactionEcommerce) (*entities.Transaction, error) {
	args := m.Called(ctx, id, transaction)
	return args.Get(0).(*entities.Transaction), args.Error(1)
}

func (m *NewTransactionRepositoryMock) Delete(ctx context.Context, id string) (*entities.Transaction, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*entities.Transaction), args.Error(1)
}

func (m *NewTransactionRepositoryMock) FindProductsByTransactionID(ctx context.Context, id string) (*aggregation.Ecommerce, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*aggregation.Ecommerce), args.Error(1)
}
