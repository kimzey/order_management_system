package repository

import (
	"context"

	_interface "github.com/kizmey/order_management_system/pkg/interface"
	"github.com/kizmey/order_management_system/pkg/interface/entities"
)

type TransactionRepository interface {
	Create(ctx context.Context, transaction *_interface.TransactionEcommerce) (*entities.Transaction, error)
	FindAll(ctx context.Context) (*[]entities.Transaction, error)

	FindByID(ctx context.Context, id string) (*entities.Transaction, error)

	Update(ctx context.Context, id string, transaction *_interface.TransactionEcommerce) (*entities.Transaction, error)
	Delete(ctx context.Context, id string) (*entities.Transaction, error)
	FindProductsByTransactionID(ctx context.Context, id string) (*_interface.Ecommerce, error)
}
