package repository

import (
	_interface "github.com/kizmey/order_management_system/pkg/interface"
	"github.com/kizmey/order_management_system/pkg/interface/entities"
)

type TransactionRepository interface {
	Create(transaction *_interface.TransactionEcommerce) (*entities.Transaction, error)
	FindAll() (*[]entities.Transaction, error)

	FindByID(id string) (*entities.Transaction, error)

	Update(id string, transaction *_interface.TransactionEcommerce) (*entities.Transaction, error)
	Delete(id string) error
	FindProductsByTransactionID(id string) (*_interface.Ecommerce, error)
}
