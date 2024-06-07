package service

import (
	"github.com/kizmey/order_management_system/entities"
	_ProductRepository "github.com/kizmey/order_management_system/pkg/product/repository"
)

type productServiceImpl struct {
	productRepository _ProductRepository.ProductRepository
}

func NewProductServiceImpl(productRepository _ProductRepository.ProductRepository) ProductService {
	return &productServiceImpl{productRepository: productRepository}
}

func (s *productServiceImpl) Create(product *entities.Product) (*entities.Product, error) {
	return s.productRepository.Create(product)
}

func (s *productServiceImpl) FindAll() (*[]entities.Product, error) {
	return s.productRepository.FindAll()
}

func (s *productServiceImpl) FindByID(id uint64) (*entities.Product, error) {
	return s.productRepository.FindByID(id)
}
func (s *productServiceImpl) Update(id uint64, product *entities.Product) (*entities.Product, error) {

	return s.productRepository.Update(id, product)
}

func (s *productServiceImpl) Delete(id uint64) error {
	return s.productRepository.Delete(id)
}
