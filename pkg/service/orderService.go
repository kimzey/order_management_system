package service

import (
	"github.com/kizmey/order_management_system/modelReq"
	"github.com/kizmey/order_management_system/modelRes"
)

type OrderService interface {
	Create(order *modelReq.Order) (*modelRes.Order, error)
	ChangeStatusNext(id uint64) (*modelRes.Order, error)
	ChageStatusDone(id uint64) (*modelRes.Order, error)
	FindAll() (*[]modelRes.Order, error)
	FindByID(id uint64) (*modelRes.Order, error)
	Update(id uint64, order *modelReq.Order) (*modelRes.Order, error)
	Delete(id uint64) error
}
