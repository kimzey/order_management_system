package order

import (
	"context"
	"errors"
	"github.com/kizmey/order_management_system/pkg/interface/aggregation"
	"testing"

	"github.com/kizmey/order_management_system/pkg/interface/entities"
	_orderRepository "github.com/kizmey/order_management_system/pkg/repository/order"
	_stockRepository "github.com/kizmey/order_management_system/pkg/repository/stock"
	_transactionRepository "github.com/kizmey/order_management_system/pkg/repository/transaction"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestOrderService_Create(t *testing.T) {
	mockOrderRepo := new(_orderRepository.NewOrderRepositoryMock)
	mockTransactionRepo := new(_transactionRepository.NewTransactionRepositoryMock)
	mockStockRepo := new(_stockRepository.NewStockRepositoryMock)
	service := NewOrderServiceImpl(mockOrderRepo, mockTransactionRepo, mockStockRepo)

	mockOrder := &entities.Order{
		OrderID:       "1",
		TransactionID: "1",
		Status:        "New",
	}

	mockEcommerce := &aggregation.Ecommerce{
		Order: &entities.Order{
			OrderID:       "1",
			TransactionID: "1",
			Status:        "New",
		},
		Product: []entities.Product{
			{
				ProductID:    "1",
				ProductName:  "Product1",
				ProductPrice: 100,
			},
			{
				ProductID:    "2",
				ProductName:  "Product2",
				ProductPrice: 200,
			},
		},
		Quantity: []uint{
			10,
			20,
		},
	}

	t.Run("successful create order", func(t *testing.T) {
		mockOrderRepo.On("Create", mock.Anything, mock.Anything).Return(mockOrder, nil).Once()
		mockTransactionRepo.On("FindProductsByTransactionID", mock.Anything, mockOrder.TransactionID).Return(mockEcommerce, nil).Once()
		mockStockRepo.On("CheckStockByProductId", mock.Anything, "1").Return(&entities.Stock{StockID: "1", ProductID: "1", Quantity: 10}, nil).Once()
		mockStockRepo.On("CheckStockByProductId", mock.Anything, "2").Return(&entities.Stock{StockID: "2", ProductID: "2", Quantity: 20}, nil).Once()
		mockStockRepo.On("Update", mock.Anything, mock.Anything, mock.Anything).Return(&entities.Stock{}, nil).Twice()

		ctx := context.Background()
		createdOrder, err := service.Create(ctx, mockOrder)

		assert.NoError(t, err)
		assert.NotNil(t, createdOrder)
		assert.Equal(t, mockOrder.OrderID, createdOrder.OrderID)
		assert.Equal(t, "New", createdOrder.Status)

		mockOrderRepo.AssertExpectations(t)
		mockTransactionRepo.AssertExpectations(t)
		mockStockRepo.AssertExpectations(t)
	})

	t.Run("create order with order repository error", func(t *testing.T) {
		mockOrderRepo.On("Create", mock.Anything, mockOrder).Return(&entities.Order{}, errors.New("order repository error")).Once()

		ctx := context.Background()
		_, err := service.Create(ctx, mockOrder)

		assert.Error(t, err)
		assert.Equal(t, "order repository error", err.Error())
		mockOrderRepo.AssertExpectations(t)
	})

	t.Run("create order with transaction repository error", func(t *testing.T) {
		mockOrderRepo.On("Create", mock.Anything, mock.Anything).Return(mockOrder, nil).Once()
		mockTransactionRepo.On("FindProductsByTransactionID", mock.Anything, mockOrder.TransactionID).Return(mockEcommerce, errors.New("transaction repository error")).Once()

		ctx := context.Background()
		_, err := service.Create(ctx, mockOrder)

		assert.Error(t, err)
		assert.Equal(t, "transaction repository error", err.Error())
		mockOrderRepo.AssertExpectations(t)
		mockTransactionRepo.AssertExpectations(t)
		mockStockRepo.AssertExpectations(t)
	})

	t.Run("create order with stock repository error and rollback", func(t *testing.T) {
		mockOrderRepo.On("Create", mock.Anything, mock.Anything).Return(mockOrder, nil).Once()
		mockTransactionRepo.On("FindProductsByTransactionID", mock.Anything, mockOrder.TransactionID).Return(mockEcommerce, nil).Once()
		mockStockRepo.On("CheckStockByProductId", mock.Anything, "1").Return(&entities.Stock{StockID: "1", ProductID: "1", Quantity: 10}, nil).Once()
		mockStockRepo.On("Update", mock.Anything, mock.Anything, mock.Anything).Return(&entities.Stock{}, errors.New("stock repository update error")).Once()
		mockOrderRepo.On("Delete", mock.Anything, mockOrder.OrderID).Return(&entities.Order{}, nil).Once()

		ctx := context.Background()
		_, err := service.Create(ctx, mockOrder)

		assert.Error(t, err)
		assert.Equal(t, "stock repository update error", err.Error())

		mockOrderRepo.AssertExpectations(t)
		mockTransactionRepo.AssertExpectations(t)
		mockStockRepo.AssertExpectations(t)

	})

	t.Run("create order with check stock error", func(t *testing.T) {
		mockOrderRepo.On("Create", mock.Anything, mock.Anything).Return(mockOrder, nil).Once()
		mockTransactionRepo.On("FindProductsByTransactionID", mock.Anything, mockOrder.TransactionID).Return(mockEcommerce, nil).Once()
		mockStockRepo.On("CheckStockByProductId", mock.Anything, "1").Return(&entities.Stock{}, errors.New("check stock error")).Once()

		ctx := context.Background()
		_, err := service.Create(ctx, mockOrder)

		assert.Error(t, err)
		assert.Equal(t, "check stock error", err.Error())
		mockOrderRepo.AssertExpectations(t)
		mockTransactionRepo.AssertExpectations(t)
		mockStockRepo.AssertExpectations(t)
	})

	t.Run("create order with Quantity is full", func(t *testing.T) {
		mockEcommerce.Quantity = nil

		mockOrderRepo.On("Create", mock.Anything, mock.Anything).Return(mockOrder, nil).Once()
		mockTransactionRepo.On("FindProductsByTransactionID", mock.Anything, mockOrder.TransactionID).Return(mockEcommerce, nil).Once()

		ctx := context.Background()
		_, err := service.Create(ctx, mockOrder)

		assert.Error(t, err)
		assert.Equal(t, "quantity is nil", err.Error())
		mockOrderRepo.AssertExpectations(t)
		mockTransactionRepo.AssertExpectations(t)
	})
}

func TestOrderService_FindAll(t *testing.T) {
	mockOrderRepo := new(_orderRepository.NewOrderRepositoryMock)
	mockTransactionRepo := new(_transactionRepository.NewTransactionRepositoryMock)
	mockStockRepo := new(_stockRepository.NewStockRepositoryMock)
	service := NewOrderServiceImpl(mockOrderRepo, mockTransactionRepo, mockStockRepo)

	mockOrders := &[]entities.Order{
		{
			OrderID:       "1",
			TransactionID: "1",
			Status:        "New",
		},
		{
			OrderID:       "2",
			TransactionID: "2",
			Status:        "Completed",
		},
	}

	t.Run("successful find all orders", func(t *testing.T) {
		mockOrderRepo.On("FindAll", mock.Anything).Return(mockOrders, nil).Once()

		ctx := context.Background()
		orders, err := service.FindAll(ctx)

		assert.NoError(t, err)
		assert.NotNil(t, orders)
		assert.Equal(t, 2, len(*orders))

		mockOrderRepo.AssertExpectations(t)
	})

	t.Run("find all orders error", func(t *testing.T) {
		mockOrderRepo.On("FindAll", mock.Anything).Return(&[]entities.Order{}, errors.New("order repository error")).Once()

		ctx := context.Background()
		orders, err := service.FindAll(ctx)

		assert.Error(t, err)
		assert.Equal(t, "order repository error", err.Error())
		assert.Nil(t, orders)
	})
}

func TestOrderService_FindByID(t *testing.T) {
	mockOrderRepo := new(_orderRepository.NewOrderRepositoryMock)
	mockTransactionRepo := new(_transactionRepository.NewTransactionRepositoryMock)
	mockStockRepo := new(_stockRepository.NewStockRepositoryMock)
	service := NewOrderServiceImpl(mockOrderRepo, mockTransactionRepo, mockStockRepo)

	mockOrder := &entities.Order{
		OrderID:       "1",
		TransactionID: "1",
		Status:        "New",
	}

	t.Run("successful find order by id", func(t *testing.T) {
		mockOrderRepo.On("FindByID", mock.Anything, mockOrder.OrderID).Return(mockOrder, nil).Once()

		ctx := context.Background()
		findOrder, err := service.FindByID(ctx, mockOrder.OrderID)

		assert.NoError(t, err)
		assert.NotNil(t, findOrder)
		assert.Equal(t, mockOrder.OrderID, findOrder.OrderID)

		mockOrderRepo.AssertExpectations(t)
	})

	t.Run("find order by id error", func(t *testing.T) {
		mockOrderRepo.On("FindByID", mock.Anything, mockOrder.OrderID).Return(&entities.Order{}, errors.New("order repository error")).Once()

		ctx := context.Background()
		findOrder, err := service.FindByID(ctx, mockOrder.OrderID)

		assert.Error(t, err)
		assert.Equal(t, "order repository error", err.Error())
		assert.Nil(t, findOrder)
	})

}

func TestOrderService_Update(t *testing.T) {
	mockOrderRepo := new(_orderRepository.NewOrderRepositoryMock)
	mockTransactionRepo := new(_transactionRepository.NewTransactionRepositoryMock)
	mockStockRepo := new(_stockRepository.NewStockRepositoryMock)
	service := NewOrderServiceImpl(mockOrderRepo, mockTransactionRepo, mockStockRepo)

	mockOrder := &entities.Order{
		OrderID:       "1",
		TransactionID: "1",
		Status:        "New",
	}

	mockEcommerce := &aggregation.Ecommerce{
		Order: &entities.Order{
			OrderID:       "1",
			TransactionID: "1",
			Status:        "New",
		},
		Product: []entities.Product{
			{
				ProductID:    "1",
				ProductName:  "Product1",
				ProductPrice: 100,
			},
			{
				ProductID:    "2",
				ProductName:  "Product2",
				ProductPrice: 200,
			},
		},
		Quantity: []uint{
			10,
			20,
		},
	}

	t.Run("successful update order", func(t *testing.T) {
		mockOrderRepo.On("Update", mock.Anything, mockOrder.OrderID, mockOrder).Return(mockOrder, nil).Once()
		mockTransactionRepo.On("FindProductsByTransactionID", mock.Anything, mockOrder.TransactionID).Return(mockEcommerce, nil).Once()
		mockStockRepo.On("CheckStockByProductId", mock.Anything, "1").Return(&entities.Stock{StockID: "1", ProductID: "1", Quantity: 10}, nil).Once()
		mockStockRepo.On("CheckStockByProductId", mock.Anything, "2").Return(&entities.Stock{StockID: "2", ProductID: "2", Quantity: 20}, nil).Once()
		mockStockRepo.On("Update", mock.Anything, mock.Anything, mock.Anything).Return(&entities.Stock{}, nil).Twice()

		ctx := context.Background()
		updatedOrder, err := service.Update(ctx, mockOrder.OrderID, mockOrder)

		assert.NoError(t, err)
		assert.NotNil(t, updatedOrder)
		assert.Equal(t, mockOrder.OrderID, updatedOrder.OrderID)
		assert.Equal(t, "New", updatedOrder.Status)

		mockOrderRepo.AssertExpectations(t)
		mockTransactionRepo.AssertExpectations(t)
		mockStockRepo.AssertExpectations(t)
	})

	t.Run("update order error", func(t *testing.T) {
		mockOrderRepo.On("Update", mock.Anything, mockOrder.OrderID, mockOrder).Return(&entities.Order{}, errors.New("order repository error")).Once()

		ctx := context.Background()
		updatedOrder, err := service.Update(ctx, mockOrder.OrderID, mockOrder)

		assert.Error(t, err)
		assert.Equal(t, "order repository error", err.Error())
		assert.Nil(t, updatedOrder)

		mockOrderRepo.AssertExpectations(t)
	})

	t.Run("update order with cheeck ProductByTransaction error", func(t *testing.T) {
		mockOrderRepo.On("Update", mock.Anything, mockOrder.OrderID, mockOrder).Return(mockOrder, nil).Once()
		mockTransactionRepo.On("FindProductsByTransactionID", mock.Anything, mockOrder.TransactionID).Return(&aggregation.Ecommerce{}, errors.New("transaction repository error")).Once()

		ctx := context.Background()
		updatedOrder, err := service.Update(ctx, mockOrder.OrderID, mockOrder)

		assert.Error(t, err)
		assert.Equal(t, "transaction repository error", err.Error())
		assert.Nil(t, updatedOrder)

		mockOrderRepo.AssertExpectations(t)
		mockTransactionRepo.AssertExpectations(t)
	})

	t.Run("update order with check Stock error", func(t *testing.T) {
		mockOrderRepo.On("Update", mock.Anything, mockOrder.OrderID, mockOrder).Return(mockOrder, nil).Once()
		mockTransactionRepo.On("FindProductsByTransactionID", mock.Anything, mockOrder.TransactionID).Return(mockEcommerce, nil).Once()
		mockStockRepo.On("CheckStockByProductId", mock.Anything, "1").Return(&entities.Stock{}, errors.New("stock repository error")).Once()

		ctx := context.Background()
		updatedOrder, err := service.Update(ctx, mockOrder.OrderID, mockOrder)

		assert.Error(t, err)
		assert.Equal(t, "stock repository error", err.Error())
		assert.Nil(t, updatedOrder)

		mockOrderRepo.AssertExpectations(t)
		mockTransactionRepo.AssertExpectations(t)
		mockStockRepo.AssertExpectations(t)
	})

	t.Run("update order with update stock error", func(t *testing.T) {
		mockOrderRepo.On("Update", mock.Anything, mockOrder.OrderID, mockOrder).Return(mockOrder, nil).Once()
		mockTransactionRepo.On("FindProductsByTransactionID", mock.Anything, mockOrder.TransactionID).Return(mockEcommerce, nil).Once()
		mockStockRepo.On("CheckStockByProductId", mock.Anything, "1").Return(&entities.Stock{StockID: "1", ProductID: "1", Quantity: 10}, nil).Once()
		mockStockRepo.On("Update", mock.Anything, mock.Anything, mock.Anything).Return(&entities.Stock{}, errors.New("stock repository error")).Once()
		mockOrderRepo.On("Delete", mock.Anything, mock.Anything).Return(&entities.Order{}, nil).Once()

		ctx := context.Background()
		updatedOrder, err := service.Update(ctx, mockOrder.OrderID, mockOrder)

		assert.Error(t, err)
		assert.Equal(t, "stock repository error", err.Error())
		assert.Nil(t, updatedOrder)

		mockOrderRepo.AssertExpectations(t)
		mockTransactionRepo.AssertExpectations(t)
		mockStockRepo.AssertExpectations(t)
	})

	t.Run("update order with quantity is nil", func(t *testing.T) {
		mockEcommerce.Quantity = nil
		mockOrderRepo.On("Update", mock.Anything, mockOrder.OrderID, mockOrder).Return(mockOrder, nil).Once()
		mockTransactionRepo.On("FindProductsByTransactionID", mock.Anything, mockOrder.TransactionID).Return(mockEcommerce, nil).Once()

		ctx := context.Background()
		updatedOrder, err := service.Update(ctx, mockOrder.OrderID, mockOrder)

		assert.Error(t, err)
		assert.Equal(t, "quantity is nil", err.Error())
		assert.Nil(t, updatedOrder)

		mockOrderRepo.AssertExpectations(t)
		mockTransactionRepo.AssertExpectations(t)
	})
}

func TestOrderService_Delete(t *testing.T) {
	mockOrderRepo := new(_orderRepository.NewOrderRepositoryMock)
	mockTransactionRepo := new(_transactionRepository.NewTransactionRepositoryMock)
	mockStockRepo := new(_stockRepository.NewStockRepositoryMock)
	service := NewOrderServiceImpl(mockOrderRepo, mockTransactionRepo, mockStockRepo)

	mockOrder := &entities.Order{
		OrderID:       "1",
		TransactionID: "1",
		Status:        "New",
	}

	t.Run("successful delete order", func(t *testing.T) {
		mockOrderRepo.On("Delete", mock.Anything, mockOrder.OrderID).Return(mockOrder, nil).Once()

		ctx := context.Background()
		deletedOrder, err := service.Delete(ctx, mockOrder.OrderID)

		assert.NoError(t, err)
		assert.NotNil(t, deletedOrder)
		assert.Equal(t, mockOrder.OrderID, deletedOrder.OrderID)

		mockOrderRepo.AssertExpectations(t)
	})

	t.Run("delete order error", func(t *testing.T) {
		mockOrderRepo.On("Delete", mock.Anything, mockOrder.OrderID).Return(&entities.Order{}, errors.New("order repository error")).Once()

		ctx := context.Background()
		deletedOrder, err := service.Delete(ctx, mockOrder.OrderID)

		assert.Error(t, err)
		assert.Equal(t, "order repository error", err.Error())
		assert.Nil(t, deletedOrder)

		mockOrderRepo.AssertExpectations(t)
	})
}

func TestOrderService_ChangeStatusNext(t *testing.T) {
	mockOrderRepo := new(_orderRepository.NewOrderRepositoryMock)
	mockTransactionRepo := new(_transactionRepository.NewTransactionRepositoryMock)
	mockStockRepo := new(_stockRepository.NewStockRepositoryMock)
	service := NewOrderServiceImpl(mockOrderRepo, mockTransactionRepo, mockStockRepo)

	mockOrder := &entities.Order{
		OrderID:       "1",
		TransactionID: "1",
		Status:        "New",
	}

	t.Run("order successful change status to next", func(t *testing.T) {
		mockOrderRepo.On("FindByID", mock.Anything, mockOrder.OrderID).Return(mockOrder, nil).Once()
		mockOrderRepo.On("UpdateStatus", mock.Anything, mockOrder.OrderID, mock.Anything).Return(mockOrder, nil).Once()

		ctx := context.Background()
		updatedOrder, err := service.ChangeStatusNext(ctx, mockOrder.OrderID)

		assert.NoError(t, err)
		assert.NotNil(t, updatedOrder)
		assert.Equal(t, mockOrder.OrderID, updatedOrder.OrderID)

		mockOrderRepo.AssertExpectations(t)
	})

	t.Run("order change status to next with FindByID error", func(t *testing.T) {
		mockOrderRepo.On("FindByID", mock.Anything, mockOrder.OrderID).Return(&entities.Order{}, errors.New("order repository error")).Once()

		ctx := context.Background()
		updatedOrder, err := service.ChangeStatusNext(ctx, mockOrder.OrderID)

		assert.Error(t, err)
		assert.Equal(t, "order repository error", err.Error())
		assert.Nil(t, updatedOrder)

		mockOrderRepo.AssertExpectations(t)
	})

	t.Run("order change status to next error", func(t *testing.T) {
		mockOrderRepo.On("FindByID", mock.Anything, mockOrder.OrderID).Return(&entities.Order{}, nil).Once()

		ctx := context.Background()
		updatedOrder, err := service.ChangeStatusNext(ctx, mockOrder.OrderID)

		assert.Error(t, err)
		assert.Equal(t, "invalid order status", err.Error())
		assert.Nil(t, updatedOrder)

		mockOrderRepo.AssertExpectations(t)
	})

	t.Run("order change updates status to next error", func(t *testing.T) {
		mockOrderRepo.On("FindByID", mock.Anything, mockOrder.OrderID).Return(mockOrder, nil).Once()
		mockOrderRepo.On("UpdateStatus", mock.Anything, mockOrder.OrderID, mock.Anything).Return(&entities.Order{}, errors.New("order repository error")).Once()

		ctx := context.Background()
		updatedOrder, err := service.ChangeStatusNext(ctx, mockOrder.OrderID)

		assert.Error(t, err)
		assert.Equal(t, "order repository error", err.Error())
		assert.Nil(t, updatedOrder)

		mockOrderRepo.AssertExpectations(t)
	})
}

func TestOrderService_ChangeStatusDone(t *testing.T) {
	mockOrderRepo := new(_orderRepository.NewOrderRepositoryMock)
	mockTransactionRepo := new(_transactionRepository.NewTransactionRepositoryMock)
	mockStockRepo := new(_stockRepository.NewStockRepositoryMock)
	service := NewOrderServiceImpl(mockOrderRepo, mockTransactionRepo, mockStockRepo)

	mockOrder := &entities.Order{
		OrderID:       "1",
		TransactionID: "1",
		Status:        "Paid",
	}

	t.Run("order successful change status to done", func(t *testing.T) {
		mockOrderRepo.On("FindByID", mock.Anything, mockOrder.OrderID).Return(mockOrder, nil).Once()
		mockOrderRepo.On("UpdateStatus", mock.Anything, mockOrder.OrderID, mock.Anything).Return(mockOrder, nil).Once()

		ctx := context.Background()
		updatedOrder, err := service.ChageStatusDone(ctx, mockOrder.OrderID)

		assert.NoError(t, err)
		assert.NotNil(t, updatedOrder)
		assert.Equal(t, mockOrder.OrderID, updatedOrder.OrderID)

		mockOrderRepo.AssertExpectations(t)
	})

	t.Run("order change status to done with FindByID error", func(t *testing.T) {
		mockOrderRepo.On("FindByID", mock.Anything, mockOrder.OrderID).Return(&entities.Order{}, errors.New("order repository error")).Once()

		ctx := context.Background()
		updatedOrder, err := service.ChageStatusDone(ctx, mockOrder.OrderID)

		assert.Error(t, err)
		assert.Equal(t, "order repository error", err.Error())
		assert.Nil(t, updatedOrder)

		mockOrderRepo.AssertExpectations(t)
	})

	t.Run("order change status to done error", func(t *testing.T) {
		mockOrderRepo.On("FindByID", mock.Anything, mockOrder.OrderID).Return(&entities.Order{}, nil).Once()

		ctx := context.Background()
		updatedOrder, err := service.ChageStatusDone(ctx, mockOrder.OrderID)

		assert.Error(t, err)
		assert.Equal(t, "invalid order status", err.Error())
		assert.Nil(t, updatedOrder)

		mockOrderRepo.AssertExpectations(t)
	})

	t.Run("order change updates status to done error", func(t *testing.T) {
		mockOrder.Status = "Paid"
		mockOrderRepo.On("FindByID", mock.Anything, mockOrder.OrderID).Return(mockOrder, nil).Once()
		mockOrderRepo.On("UpdateStatus", mock.Anything, mockOrder.OrderID, mock.Anything).Return(&entities.Order{}, errors.New("order repository error")).Once()

		ctx := context.Background()
		updatedOrder, err := service.ChageStatusDone(ctx, mockOrder.OrderID)

		assert.Error(t, err)
		assert.Equal(t, "order repository error", err.Error())
		assert.Nil(t, updatedOrder)

		mockOrderRepo.AssertExpectations(t)
	})

}
