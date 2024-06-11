package repository

import (
	"github.com/kizmey/order_management_system/pkg/interface/entities"
)

type ProductRepository interface {
	Create(product *entities.Product) (*entities.Product, error)
	FindAll() (*[]entities.Product, error)
	FindByID(id string) (*entities.Product, error)
	Update(id string, product *entities.Product) (*entities.Product, error)
	Delete(id string) error
}
