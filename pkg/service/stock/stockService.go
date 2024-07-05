package stock

import (
	"context"
	"github.com/kizmey/order_management_system/pkg/interface/entities"
	"go.opentelemetry.io/otel"
)

type StockService interface {
	Create(ctx context.Context, stock *entities.Stock) (*entities.Stock, error)
	FindAll(ctx context.Context) (*[]entities.Stock, error)
	CheckStockByProductId(ctx context.Context, id string) (*entities.Stock, error)
	Update(ctx context.Context, id string, stock *entities.Stock) (*entities.Stock, error)
	Delete(ctx context.Context, id string) (*entities.Stock, error)
}

var tracer = otel.Tracer("StockService")
