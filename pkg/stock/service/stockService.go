package service

import "github.com/kizmey/order_management_system/entities"

type StockService interface {
	Create(stock *entities.Stock) (*entities.Stock, error)
	FindAll() (*[]entities.Stock, error)
	CheckStockByProductId(id uint64) (*entities.Stock, error)
	Update(id uint64, stock *entities.Stock) (*entities.Stock, error)
	Delete(id uint64) error
}
