package stock

import (
	"context"
	"github.com/kizmey/order_management_system/pkg/interface/entities"
	"github.com/stretchr/testify/mock"
)

type NewStockRepositoryMock struct {
	mock.Mock
}

func (m *NewStockRepositoryMock) checkStockRepositoryMock() StockRepository {
	return m
}

func (m *NewStockRepositoryMock) Create(ctx context.Context, stock *entities.Stock) (*entities.Stock, error) {
	args := m.Called(ctx, stock)
	return args.Get(0).(*entities.Stock), args.Error(1)
}

func (m *NewStockRepositoryMock) FindAll(ctx context.Context) (*[]entities.Stock, error) {
	args := m.Called(ctx)
	return args.Get(0).(*[]entities.Stock), args.Error(1)
}

func (m *NewStockRepositoryMock) CheckStockByProductId(ctx context.Context, id string) (*entities.Stock, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*entities.Stock), args.Error(1)
}

func (m *NewStockRepositoryMock) Update(ctx context.Context, id string, stock *entities.Stock) (*entities.Stock, error) {
	args := m.Called(ctx, id, stock)
	return args.Get(0).(*entities.Stock), args.Error(1)
}

func (m *NewStockRepositoryMock) Delete(ctx context.Context, id string) (*entities.Stock, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*entities.Stock), args.Error(1)
}
