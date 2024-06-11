package service

import (
	"github.com/kizmey/order_management_system/pkg/interface/modelReq"
	"github.com/kizmey/order_management_system/pkg/interface/modelRes"
)

type TransactionService interface {
	Create(transaction *modelReq.Transaction) (*modelRes.Transaction, error)
	FindAll() (*[]modelRes.Transaction, error)
	FindByID(id string) (*modelRes.Transaction, error)
	Update(id string, transaction *modelReq.Transaction) (*modelRes.Transaction, error)
	Delete(id string) error
}
