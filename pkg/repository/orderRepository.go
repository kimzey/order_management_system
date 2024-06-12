package repository

import (
	_interface "github.com/kizmey/order_management_system/pkg/interface"
	"github.com/kizmey/order_management_system/pkg/interface/entities"
)

type OrderRepository interface {
	Create(order *_interface.Ecommerce) (*entities.Order, error)
	FindAll() (*[]entities.Order, error)
	FindByID(id string) (*entities.Order, error)
	Update(id string, order *_interface.Ecommerce) (*entities.Order, error)
	UpdateStatus(id string, order *entities.Order) (*entities.Order, error)
	Delete(id string) error
}
