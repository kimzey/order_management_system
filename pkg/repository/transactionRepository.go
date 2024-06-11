package repository

import (
	"github.com/kizmey/order_management_system/pkg/interface/entities"
)

type TransactionRepository interface {
	Create(transaction *entities.Transaction) (*entities.Transaction, error)
	FindAll() (*[]entities.Transaction, error)

	FindByID(id uint64) (*entities.Transaction, error)

	Update(id uint64, transaction *entities.Transaction) (*entities.Transaction, error)
	Delete(id uint64) error
}
