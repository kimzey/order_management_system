package product

import (
	"context"
	"go.opentelemetry.io/otel"

	"github.com/kizmey/order_management_system/pkg/interface/entities"
)

type ProductRepository interface {
	Create(ctx context.Context, product *entities.Product) (*entities.Product, error)
	FindAll(ctx context.Context) (*[]entities.Product, error)
	FindByID(ctx context.Context, id string) (*entities.Product, error)
	Update(ctx context.Context, id string, product *entities.Product) (*entities.Product, error)
	Delete(ctx context.Context, id string) (*entities.Product, error)
}

var tracer = otel.Tracer("ProductRepository")
