package repository

import (
	"github.com/kizmey/order_management_system/pkg/interface/entities"
)

type TransactionRepository interface {
	Create(transaction *entities.Transaction) (*entities.Transaction, error)
	FindAll() (*[]entities.Transaction, error)

	FindByID(id string) (*entities.Transaction, error)

	Update(id string, transaction *entities.Transaction) (*entities.Transaction, error)
	Delete(id string) error
}
