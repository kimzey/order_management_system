package service

import (
	"context"
	"github.com/kizmey/order_management_system/pkg/interface/modelReq"
	"github.com/kizmey/order_management_system/pkg/interface/modelRes"
)

type ProductService interface {
	Create(ctx context.Context, product *modelReq.Product) (*modelRes.Product, error)
	FindAll(ctx context.Context) (*[]modelRes.Product, error)
	FindByID(ctx context.Context, id string) (*modelRes.Product, error)
	Update(ctx context.Context, id string, product *modelReq.Product) (*modelRes.Product, error)
	Delete(ctx context.Context, id string) (*modelRes.Product, error)
}
