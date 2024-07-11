package test

import (
	"context"
	"errors"
	"testing"

	"github.com/kizmey/order_management_system/pkg/interface/entities"
	_stockRepository "github.com/kizmey/order_management_system/pkg/repository/stock"
	_stockService "github.com/kizmey/order_management_system/pkg/service/stock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestStockService_Create(t *testing.T) {
	mockRepo := new(_stockRepository.NewStockRepositoryMock)
	service := _stockService.NewStockServiceImpl(mockRepo)

	stock := &entities.Stock{StockID: "1", ProductID: "1", Quantity: 10}

	t.Run("successful create", func(t *testing.T) {
		mockRepo.On("Create", mock.Anything, stock).Return(stock, nil).Once()

		ctx := context.Background()
		createdStock, err := service.Create(ctx, stock)

		assert.NoError(t, err)
		assert.Equal(t, stock, createdStock)
		mockRepo.AssertExpectations(t)
	})

	t.Run("create with error", func(t *testing.T) {
		mockRepo.On("Create", mock.Anything, stock).Return(&entities.Stock{}, errors.New("create error")).Once()

		ctx := context.Background()
		createdStock, err := service.Create(ctx, stock)

		assert.Error(t, err)
		assert.Nil(t, createdStock)
		assert.Equal(t, "create error", err.Error())
		mockRepo.AssertExpectations(t)
	})
}

func TestStockService_FindAll(t *testing.T) {
	mockRepo := new(_stockRepository.NewStockRepositoryMock)
	service := _stockService.NewStockServiceImpl(mockRepo)

	mockStocks := []entities.Stock{
		{StockID: "1", ProductID: "1", Quantity: 10},
		{StockID: "2", ProductID: "2", Quantity: 20},
	}

	t.Run("successful find all", func(t *testing.T) {
		mockRepo.On("FindAll", mock.Anything).Return(&mockStocks, nil).Once()

		ctx := context.Background()
		foundStocks, err := service.FindAll(ctx)

		assert.NoError(t, err)
		assert.NotNil(t, foundStocks)
		assert.Equal(t, len(mockStocks), len(*foundStocks))
		mockRepo.AssertExpectations(t)
	})

	t.Run("find all with error", func(t *testing.T) {
		// Return an empty slice instead of nil
		mockRepo.On("FindAll", mock.Anything).Return(&[]entities.Stock{}, errors.New("find all error")).Once()

		ctx := context.Background()
		foundStocks, err := service.FindAll(ctx)

		assert.Error(t, err)
		assert.Nil(t, foundStocks)
		assert.Equal(t, "find all error", err.Error())
		mockRepo.AssertExpectations(t)
	})
}
func TestStockService_CheckStockByProductId(t *testing.T) {
	mockRepo := new(_stockRepository.NewStockRepositoryMock)
	service := _stockService.NewStockServiceImpl(mockRepo)

	mockStock := &entities.Stock{StockID: "1", ProductID: "1", Quantity: 10}

	t.Run("successful check stock by product ID", func(t *testing.T) {
		mockRepo.On("CheckStockByProductId", mock.Anything, "1").Return(mockStock, nil).Once()

		ctx := context.Background()
		foundStock, err := service.CheckStockByProductId(ctx, "1")

		assert.NoError(t, err)
		assert.NotNil(t, foundStock)
		assert.Equal(t, mockStock.StockID, foundStock.StockID)
		assert.Equal(t, mockStock.ProductID, foundStock.ProductID)
		assert.Equal(t, mockStock.Quantity, foundStock.Quantity)
		mockRepo.AssertExpectations(t)
	})

	t.Run("check stock by product ID with error", func(t *testing.T) {
		mockRepo.On("CheckStockByProductId", mock.Anything, "1").Return(&entities.Stock{}, errors.New("check stock error")).Once()

		ctx := context.Background()
		foundStock, err := service.CheckStockByProductId(ctx, "1")

		assert.Error(t, err)
		assert.Nil(t, foundStock)
		assert.Equal(t, "check stock error", err.Error())
		mockRepo.AssertExpectations(t)
	})
}

func TestStockService_Update(t *testing.T) {
	mockRepo := new(_stockRepository.NewStockRepositoryMock)
	service := _stockService.NewStockServiceImpl(mockRepo)

	mockStock := &entities.Stock{StockID: "1", ProductID: "1", Quantity: 10}

	t.Run("successful update", func(t *testing.T) {
		mockRepo.On("Update", mock.Anything, "1", mockStock).Return(mockStock, nil).Once()

		ctx := context.Background()
		updatedStock, err := service.Update(ctx, "1", mockStock)

		assert.NoError(t, err)
		assert.NotNil(t, updatedStock)
		assert.Equal(t, mockStock.StockID, updatedStock.StockID)
		assert.Equal(t, mockStock.ProductID, updatedStock.ProductID)
		assert.Equal(t, mockStock.Quantity, updatedStock.Quantity)
		mockRepo.AssertExpectations(t)
	})

	t.Run("update with error", func(t *testing.T) {
		mockRepo.On("Update", mock.Anything, "1", mockStock).Return(&entities.Stock{}, errors.New("update error")).Once()

		ctx := context.Background()
		updatedStock, err := service.Update(ctx, "1", mockStock)

		assert.Error(t, err)
		assert.Nil(t, updatedStock)
		assert.Equal(t, "update error", err.Error())
		mockRepo.AssertExpectations(t)
	})
}

func TestStockService_Delete(t *testing.T) {
	mockRepo := new(_stockRepository.NewStockRepositoryMock)
	service := _stockService.NewStockServiceImpl(mockRepo)

	mockStock := &entities.Stock{StockID: "1", ProductID: "1", Quantity: 10}

	t.Run("successful delete", func(t *testing.T) {
		mockRepo.On("Delete", mock.Anything, "1").Return(mockStock, nil).Once()

		ctx := context.Background()
		deletedStock, err := service.Delete(ctx, "1")

		assert.NoError(t, err)
		assert.NotNil(t, deletedStock)
		assert.Equal(t, mockStock.StockID, deletedStock.StockID)
		assert.Equal(t, mockStock.ProductID, deletedStock.ProductID)
		assert.Equal(t, mockStock.Quantity, deletedStock.Quantity)
		mockRepo.AssertExpectations(t)
	})

	t.Run("delete with error", func(t *testing.T) {
		mockRepo.On("Delete", mock.Anything, "1").Return(&entities.Stock{}, errors.New("delete error")).Once()

		ctx := context.Background()
		deletedStock, err := service.Delete(ctx, "1")

		assert.Error(t, err)
		assert.Nil(t, deletedStock)
		assert.Equal(t, "delete error", err.Error())
		mockRepo.AssertExpectations(t)
	})
}
