package repository

import (
	"github.com/kizmey/order_management_system/pkg/interface/entities"
)

type StockRepository interface {
	Create(stock *entities.Stock) (*entities.Stock, error)
	FindAll() (*[]entities.Stock, error)
	CheckStockByProductId(id string) (*entities.Stock, error)
	Update(id string, stock *entities.Stock) (*entities.Stock, error)
	Delete(id string) error
}
