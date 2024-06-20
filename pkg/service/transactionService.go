package service

import (
	"context"
	"github.com/kizmey/order_management_system/pkg/interface/modelReq"
	"github.com/kizmey/order_management_system/pkg/interface/modelRes"
)

type TransactionService interface {
	Create(ctx context.Context, transaction *modelReq.Transaction) (*modelRes.Transaction, error)
	FindAll(ctx context.Context) (*[]modelRes.Transaction, error)
	FindByID(ctx context.Context, id string) (*modelRes.Transaction, error)
	Update(ctx context.Context, id string, transaction *modelReq.Transaction) (*modelRes.Transaction, error)
	Delete(ctx context.Context, id string) (*modelRes.Transaction, error)
}
