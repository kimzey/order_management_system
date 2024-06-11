package repository

import (
	"github.com/kizmey/order_management_system/pkg/interface/entities"
)

type ProductRepository interface {
	Create(product *entities.Product) (*entities.Product, error)
	FindAll() (*[]entities.Product, error)
	FindByID(id uint64) (*entities.Product, error)
	Update(id uint64, product *entities.Product) (*entities.Product, error)
	Delete(id uint64) error
}
