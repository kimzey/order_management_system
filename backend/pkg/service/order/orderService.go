package order

import (
	"context"
	"github.com/kizmey/order_management_system/pkg/interface/entities"
	"go.opentelemetry.io/otel"
)

type OrderService interface {
	Create(ctx context.Context, order *entities.Order) (*entities.Order, error)
	ChangeStatusNext(ctx context.Context, id string) (*entities.Order, error)
	ChageStatusDone(ctx context.Context, id string) (*entities.Order, error)
	FindAll(ctx context.Context) (*[]entities.Order, error)
	FindByID(ctx context.Context, id string) (*entities.Order, error)
	Update(ctx context.Context, id string, order *entities.Order) (*entities.Order, error)
	Delete(ctx context.Context, id string) (*entities.Order, error)
}

var tracer = otel.Tracer("OrderService")
