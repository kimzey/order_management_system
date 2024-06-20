package repository

import (
	"context"
	"github.com/kizmey/order_management_system/pkg/interface/entities"
)

type OrderRepository interface {
	Create(ctx context.Context, order *entities.Order) (*entities.Order, error)
	FindAll(ctx context.Context) (*[]entities.Order, error)
	FindByID(ctx context.Context, id string) (*entities.Order, error)
	Update(ctx context.Context, id string, order *entities.Order) (*entities.Order, error)
	UpdateStatus(ctx context.Context, id string, order *entities.Order) (*entities.Order, error)
	Delete(ctx context.Context, id string) (*entities.Order, error)
}
