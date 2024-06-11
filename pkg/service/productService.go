package service

import (
	"github.com/kizmey/order_management_system/pkg/interface/modelReq"
	"github.com/kizmey/order_management_system/pkg/interface/modelRes"
)

type ProductService interface {
	Create(product *modelReq.Product) (*modelRes.Product, error)
	FindAll() (*[]modelRes.Product, error)
	FindByID(id string) (*modelRes.Product, error)
	Update(id string, product *modelReq.Product) (*modelRes.Product, error)
	Delete(id string) error
}
