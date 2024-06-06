package repository

import "github.com/kizmey/order_management_system/entities"

type OrderRepository interface {
	Create(order *entities.Order) (*entities.Order, error)
	FindAll() (*[]entities.Order, error)
	//FindByID(id uint64) (*entities.Order, error)
	//Update(order *entities.Order) (*entities.Order, error)
	//Delete(order *entities.Order)
	ChangeStatusNext(id uint64) (*entities.Order, error)
	ChageStatusDone(id uint64) (*entities.Order, error)
}
