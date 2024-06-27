package transaction

import (
	"context"
	"github.com/kizmey/order_management_system/pkg/interface/entities"
	"go.opentelemetry.io/otel"
)

type TransactionService interface {
	Create(ctx context.Context, transaction *entities.Transaction) (*entities.Transaction, error)
	FindAll(ctx context.Context) (*[]entities.Transaction, error)
	FindByID(ctx context.Context, id string) (*entities.Transaction, error)
	Update(ctx context.Context, id string, transaction *entities.Transaction) (*entities.Transaction, error)
	Delete(ctx context.Context, id string) (*entities.Transaction, error)
}

var tracer = otel.Tracer("TransactionService")
