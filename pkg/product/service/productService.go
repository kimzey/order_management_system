package service

import (
	"github.com/kizmey/order_management_system/pkg/modelReq"
	"github.com/kizmey/order_management_system/pkg/modelRes"
)

type ProductService interface {
	Create(product *modelReq.Product) (*modelRes.Product, error)
	FindAll() (*[]modelRes.Product, error)
	FindByID(id uint64) (*modelRes.Product, error)
	Update(id uint64, product *modelReq.Product) (*modelRes.Product, error)
	Delete(id uint64) error
}
