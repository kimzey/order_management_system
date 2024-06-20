package service

import (
	"context"
	"github.com/kizmey/order_management_system/pkg/interface/modelReq"
	"github.com/kizmey/order_management_system/pkg/interface/modelRes"
)

type OrderService interface {
	Create(ctx context.Context, order *modelReq.Order) (*modelRes.Order, error)
	ChangeStatusNext(ctx context.Context, id string) (*modelRes.Order, error)
	ChageStatusDone(ctx context.Context, id string) (*modelRes.Order, error)
	FindAll(ctx context.Context) (*[]modelRes.Order, error)
	FindByID(ctx context.Context, id string) (*modelRes.Order, error)
	Update(ctx context.Context, id string, order *modelReq.Order) (*modelRes.Order, error)
	Delete(ctx context.Context, id string) (*modelRes.Order, error)
}
