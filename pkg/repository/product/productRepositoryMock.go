package product

import (
	"context"
	"github.com/kizmey/order_management_system/pkg/interface/entities"
	"github.com/stretchr/testify/mock"
)

type NewProductRepositoryMock struct {
	mock.Mock
}

func (m *NewProductRepositoryMock) CheckProductRepository() ProductRepository {
	return m
}

func (m *NewProductRepositoryMock) Create(ctx context.Context, product *entities.Product) (*entities.Product, error) {
	args := m.Called(ctx, product)
	if args.Get(0) != nil {
		return args.Get(0).(*entities.Product), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *NewProductRepositoryMock) FindAll(ctx context.Context) (*[]entities.Product, error) {
	args := m.Called(ctx)
	if args.Get(0) != nil {
		return args.Get(0).(*[]entities.Product), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *NewProductRepositoryMock) FindByID(ctx context.Context, id string) (*entities.Product, error) {
	args := m.Called(ctx, id)
	if args.Get(0) != nil {
		return args.Get(0).(*entities.Product), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *NewProductRepositoryMock) Update(ctx context.Context, id string, product *entities.Product) (*entities.Product, error) {
	args := m.Called(ctx, id, product)
	if args.Get(0) != nil {
		return args.Get(0).(*entities.Product), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *NewProductRepositoryMock) Delete(ctx context.Context, id string) (*entities.Product, error) {
	args := m.Called(ctx, id)
	if args.Get(0) != nil {
		return args.Get(0).(*entities.Product), args.Error(1)
	}
	return nil, args.Error(1)
}
