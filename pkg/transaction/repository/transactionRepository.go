package repository

import "github.com/kizmey/order_management_system/entities"

type TransactionRepository interface {
	Create(transaction *entities.Transaction) (uint64, error)
	FindAll() (*[]entities.Transaction, error)
}
