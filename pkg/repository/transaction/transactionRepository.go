package transaction

import (
	"context"
	"github.com/kizmey/order_management_system/pkg/interface/aggregation"
	"go.opentelemetry.io/otel"

	"github.com/kizmey/order_management_system/pkg/interface/entities"
)

type TransactionRepository interface {
	Create(ctx context.Context, transaction *aggregation.TransactionEcommerce) (*entities.Transaction, error)
	FindAll(ctx context.Context) (*[]entities.Transaction, error)

	FindByID(ctx context.Context, id string) (*entities.Transaction, error)

	Update(ctx context.Context, id string, transaction *aggregation.TransactionEcommerce) (*entities.Transaction, error)
	Delete(ctx context.Context, id string) (*entities.Transaction, error)
	FindProductsByTransactionID(ctx context.Context, id string) (*aggregation.Ecommerce, error)
}

var tracer = otel.Tracer("TransactionRepository")
