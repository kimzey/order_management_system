package service

import "github.com/kizmey/order_management_system/entities"

type TransactionService interface {
	Create(transaction *entities.Transaction) (uint64, error)
	FindAll() (*[]entities.Transaction, error)
	FindByID(id uint64) (*entities.Transaction, error)
	Update(id uint64, transaction *entities.Transaction) (*entities.Transaction, error)
	Delete(id uint64) error
}
