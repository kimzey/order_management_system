package transaction

import (
	"context"
	_interface "github.com/kizmey/order_management_system/pkg/interface/aggregation"
	"github.com/kizmey/order_management_system/pkg/interface/entities"
	"go.opentelemetry.io/otel"
)

type TransactionService interface {
	Create(ctx context.Context, transaction *_interface.TransactionEcommerce) (*_interface.TransactionEcommerce, error)
	FindAll(ctx context.Context) (*[]entities.Transaction, error)
	FindByID(ctx context.Context, id string) (*entities.Transaction, error)
	Update(ctx context.Context, id string, transaction *_interface.TransactionEcommerce) (*_interface.TransactionEcommerce, error)
	Delete(ctx context.Context, id string) (*entities.Transaction, error)
}

var tracer = otel.Tracer("TransactionService")
