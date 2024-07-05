package order

import (
	"context"
	"github.com/kizmey/order_management_system/pkg/interface/entities"
	"github.com/stretchr/testify/mock"
)

type NewOrderRepositoryMock struct {
	mock.Mock
}

func (m *NewOrderRepositoryMock) checkOrderRepositoryMock() OrderRepository {
	return m
}

func (m *NewOrderRepositoryMock) Create(ctx context.Context, order *entities.Order) (*entities.Order, error) {
	args := m.Called(ctx, order)
	return args.Get(0).(*entities.Order), args.Error(1)
}

func (m *NewOrderRepositoryMock) FindAll(ctx context.Context) (*[]entities.Order, error) {
	args := m.Called(ctx)
	return args.Get(0).(*[]entities.Order), args.Error(1)
}

func (m *NewOrderRepositoryMock) FindByID(ctx context.Context, id string) (*entities.Order, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*entities.Order), args.Error(1)
}

func (m *NewOrderRepositoryMock) Update(ctx context.Context, id string, order *entities.Order) (*entities.Order, error) {
	args := m.Called(ctx, id, order)
	return args.Get(0).(*entities.Order), args.Error(1)
}

func (m *NewOrderRepositoryMock) UpdateStatus(ctx context.Context, id string, order *entities.Order) (*entities.Order, error) {
	args := m.Called(ctx, id, order)
	return args.Get(0).(*entities.Order), args.Error(1)
}

func (m *NewOrderRepositoryMock) Delete(ctx context.Context, id string) (*entities.Order, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*entities.Order), args.Error(1)
}
