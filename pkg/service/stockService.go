package service

import (
	"github.com/kizmey/order_management_system/pkg/interface/modelReq"
	"github.com/kizmey/order_management_system/pkg/interface/modelRes"
)

type StockService interface {
	Create(stock *modelReq.Stock) (*modelRes.Stock, error)
	FindAll() (*[]modelRes.Stock, error)
	CheckStockByProductId(id string) (*modelRes.Stock, error)
	Update(id string, stock *modelReq.Stock) (*modelRes.Stock, error)
	Delete(id string) (*modelRes.Stock, error)
}
