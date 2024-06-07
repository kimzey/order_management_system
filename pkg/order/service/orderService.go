package service

import "github.com/kizmey/order_management_system/entities"

type OrderService interface {
	Create(order *entities.Order) (*entities.Order, error)
	ChangeStatusNext(id uint64) (*entities.Order, error)
	ChageStatusDone(id uint64) (*entities.Order, error)
	FindAll() (*[]entities.Order, error)
	FindByID(id uint64) (*entities.Order, error)
	Update(id uint64, order *entities.Order) (*entities.Order, error)
	Delete(id uint64) error
}
