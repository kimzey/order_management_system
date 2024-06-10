package service

import (
	"github.com/kizmey/order_management_system/pkg/modelReq"
	"github.com/kizmey/order_management_system/pkg/modelRes"
)

type StockService interface {
	Create(stock *modelReq.Stock) (*modelRes.Stock, error)
	FindAll() (*[]modelRes.Stock, error)
	CheckStockByProductId(id uint64) (*modelRes.Stock, error)
	Update(id uint64, stock *modelReq.Stock) (*modelRes.Stock, error)
	Delete(id uint64) error
}
