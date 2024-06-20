package service

import (
	"context"
	"github.com/kizmey/order_management_system/pkg/interface/modelReq"
	"github.com/kizmey/order_management_system/pkg/interface/modelRes"
)

type StockService interface {
	Create(ctx context.Context, stock *modelReq.Stock) (*modelRes.Stock, error)
	FindAll(ctx context.Context) (*[]modelRes.Stock, error)
	CheckStockByProductId(ctx context.Context, id string) (*modelRes.Stock, error)
	Update(ctx context.Context, id string, stock *modelReq.Stock) (*modelRes.Stock, error)
	Delete(ctx context.Context, id string) (*modelRes.Stock, error)
}
