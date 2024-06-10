package service

import (
	"github.com/kizmey/order_management_system/pkg/modelReq"
	"github.com/kizmey/order_management_system/pkg/modelRes"
)

type TransactionService interface {
	Create(transaction *modelReq.Transaction) (*modelRes.Transaction, error)
	FindAll() (*[]modelRes.Transaction, error)
	FindByID(id uint64) (*modelRes.Transaction, error)
	Update(id uint64, transaction *modelReq.Transaction) (*modelRes.Transaction, error)
	Delete(id uint64) error
}
