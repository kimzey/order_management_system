package test

import (
	"context"
	"errors"
	"testing"

	"github.com/kizmey/order_management_system/pkg/interface/entities"
	_productRepository "github.com/kizmey/order_management_system/pkg/repository/product"
	_productService "github.com/kizmey/order_management_system/pkg/service/product"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestProductService_Create(t *testing.T) {
	mockRepo := new(_productRepository.NewProductRepositoryMock)
	service := _productService.NewProductServiceImpl(mockRepo)

	product := &entities.Product{ProductID: "1", ProductName: "Product1", ProductPrice: 100}

	t.Run("successful create", func(t *testing.T) {
		mockRepo.On("Create", mock.Anything, product).Return(product, nil).Once()

		ctx := context.Background()
		createdProduct, err := service.Create(ctx, product)

		assert.NoError(t, err)
		assert.Equal(t, product, createdProduct)
		mockRepo.AssertExpectations(t)
	})

	t.Run("create with error", func(t *testing.T) {
		mockRepo.On("Create", mock.Anything, product).Return(nil, errors.New("create error")).Once()

		ctx := context.Background()
		createdProduct, err := service.Create(ctx, product)

		assert.Error(t, err)
		assert.Nil(t, createdProduct)
		assert.Equal(t, "create error", err.Error())
		mockRepo.AssertExpectations(t)
	})
}

func TestProductService_FindAll(t *testing.T) {
	mockRepo := new(_productRepository.NewProductRepositoryMock)
	service := _productService.NewProductServiceImpl(mockRepo)

	// Prepare mock data
	mockProducts := []entities.Product{
		{ProductID: "1", ProductName: "Product1", ProductPrice: 100},
		{ProductID: "2", ProductName: "Product2", ProductPrice: 200},
	}

	t.Run("successful find all", func(t *testing.T) {
		mockRepo.On("FindAll", mock.Anything).Return(&mockProducts, nil).Once()

		ctx := context.Background()
		foundProducts, err := service.FindAll(ctx)

		assert.NoError(t, err)
		assert.NotNil(t, foundProducts)
		assert.Equal(t, len(mockProducts), len(*foundProducts))
		mockRepo.AssertExpectations(t)
	})

	t.Run("find all with error", func(t *testing.T) {
		mockRepo.On("FindAll", mock.Anything).Return(nil, errors.New("find all error")).Once()

		ctx := context.Background()
		foundProducts, err := service.FindAll(ctx)

		assert.Error(t, err)
		assert.Nil(t, foundProducts)
		assert.Equal(t, "find all error", err.Error())
		mockRepo.AssertExpectations(t)
	})
}

func TestProductService_FindByID(t *testing.T) {
	mockRepo := new(_productRepository.NewProductRepositoryMock)
	service := _productService.NewProductServiceImpl(mockRepo)

	// Prepare mock data
	mockProduct := &entities.Product{ProductID: "1", ProductName: "Product1", ProductPrice: 100}

	t.Run("successful find by id", func(t *testing.T) {
		mockRepo.On("FindByID", mock.Anything, "1").Return(mockProduct, nil).Once()

		ctx := context.Background()
		foundProduct, err := service.FindByID(ctx, "1")

		assert.NoError(t, err)
		assert.NotNil(t, foundProduct)
		assert.Equal(t, mockProduct.ProductID, foundProduct.ProductID)
		assert.Equal(t, mockProduct.ProductName, foundProduct.ProductName)
		assert.Equal(t, mockProduct.ProductPrice, foundProduct.ProductPrice)
		mockRepo.AssertExpectations(t)
	})

	t.Run("find by id with error", func(t *testing.T) {
		mockRepo.On("FindByID", mock.Anything, "1").Return(nil, errors.New("find by id error")).Once()

		ctx := context.Background()
		foundProduct, err := service.FindByID(ctx, "1")

		assert.Error(t, err)
		assert.Nil(t, foundProduct)
		assert.Equal(t, "find by id error", err.Error())
		mockRepo.AssertExpectations(t)
	})
}

func TestProductService_Update(t *testing.T) {
	mockRepo := new(_productRepository.NewProductRepositoryMock)
	service := _productService.NewProductServiceImpl(mockRepo)

	// Prepare mock data
	mockProduct := &entities.Product{ProductID: "1", ProductName: "UpdatedProduct", ProductPrice: 150}

	t.Run("successful update", func(t *testing.T) {
		mockRepo.On("Update", mock.Anything, "1", mockProduct).Return(mockProduct, nil).Once()

		ctx := context.Background()
		updatedProduct, err := service.Update(ctx, "1", mockProduct)

		assert.NoError(t, err)
		assert.NotNil(t, updatedProduct)
		assert.Equal(t, mockProduct.ProductName, updatedProduct.ProductName)
		assert.Equal(t, mockProduct.ProductPrice, updatedProduct.ProductPrice)
		mockRepo.AssertExpectations(t)
	})

	t.Run("update with error", func(t *testing.T) {
		mockRepo.On("Update", mock.Anything, "1", mockProduct).Return(nil, errors.New("update error")).Once()

		ctx := context.Background()
		updatedProduct, err := service.Update(ctx, "1", mockProduct)

		assert.Error(t, err)
		assert.Nil(t, updatedProduct)
		assert.Equal(t, "update error", err.Error())
		mockRepo.AssertExpectations(t)
	})
}

func TestProductService_Delete(t *testing.T) {
	mockRepo := new(_productRepository.NewProductRepositoryMock)
	service := _productService.NewProductServiceImpl(mockRepo)

	// Prepare mock data
	mockProduct := &entities.Product{ProductID: "1", ProductName: "Product1", ProductPrice: 100}

	t.Run("successful delete", func(t *testing.T) {
		mockRepo.On("Delete", mock.Anything, "1").Return(mockProduct, nil).Once()

		ctx := context.Background()
		deletedProduct, err := service.Delete(ctx, "1")

		assert.NoError(t, err)
		assert.NotNil(t, deletedProduct)
		assert.Equal(t, mockProduct.ProductID, deletedProduct.ProductID)
		assert.Equal(t, mockProduct.ProductName, deletedProduct.ProductName)
		assert.Equal(t, mockProduct.ProductPrice, deletedProduct.ProductPrice)
		mockRepo.AssertExpectations(t)
	})

	t.Run("delete with error", func(t *testing.T) {
		mockRepo.On("Delete", mock.Anything, "1").Return(nil, errors.New("delete error")).Once()

		ctx := context.Background()
		deletedProduct, err := service.Delete(ctx, "1")

		assert.Error(t, err)
		assert.Nil(t, deletedProduct)
		assert.Equal(t, "delete error", err.Error())
		mockRepo.AssertExpectations(t)
	})
}
