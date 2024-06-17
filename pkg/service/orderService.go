package service

import (
	"github.com/kizmey/order_management_system/pkg/interface/modelReq"
	"github.com/kizmey/order_management_system/pkg/interface/modelRes"
)

type OrderService interface {
	Create(order *modelReq.Order) (*modelRes.Order, error)
	ChangeStatusNext(id string) (*modelRes.Order, error)
	ChageStatusDone(id string) (*modelRes.Order, error)
	FindAll() (*[]modelRes.Order, error)
	FindByID(id string) (*modelRes.Order, error)
	Update(id string, order *modelReq.Order) (*modelRes.Order, error)
	Delete(id string) (*modelRes.Order, error)
}
