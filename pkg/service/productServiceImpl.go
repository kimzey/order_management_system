package service

import (
	"github.com/kizmey/order_management_system/pkg/interface/entities"
	"github.com/kizmey/order_management_system/pkg/interface/modelReq"
	"github.com/kizmey/order_management_system/pkg/interface/modelRes"
	_ProductRepository "github.com/kizmey/order_management_system/pkg/repository"
)

type productServiceImpl struct {
	productRepository _ProductRepository.ProductRepository
}

func NewProductServiceImpl(productRepository _ProductRepository.ProductRepository) ProductService {
	return &productServiceImpl{productRepository: productRepository}
}

func (s *productServiceImpl) Create(product *modelReq.Product) (*modelRes.Product, error) {

	productEntity := s.productReqToEntity(product)

	productEntity, err := s.productRepository.Create(productEntity)
	if err != nil {
		return nil, err
	}
	return s.productEntityToRes(productEntity), nil
}

func (s *productServiceImpl) FindAll() (*[]modelRes.Product, error) {

	products, err := s.productRepository.FindAll()
	if err != nil {
		return nil, err
	}

	productsRes := make([]modelRes.Product, 0)
	for _, product := range *products {
		productsRes = append(productsRes, *s.productEntityToRes(&product))
	}
	return &productsRes, nil
}

func (s *productServiceImpl) FindByID(id string) (*modelRes.Product, error) {

	product, err := s.productRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return s.productEntityToRes(product), nil
}
func (s *productServiceImpl) Update(id string, product *modelReq.Product) (*modelRes.Product, error) {

	productEntity := s.productReqToEntity(product)
	productEntity, err := s.productRepository.Update(id, productEntity)
	if err != nil {
		return nil, err
	}
	return s.productEntityToRes(productEntity), nil
}

func (s *productServiceImpl) Delete(id string) (*modelRes.Product, error) {

	product, err := s.productRepository.Delete(id)
	if err != nil {
		return nil, err
	}
	return s.productEntityToRes(product), nil
}

func (s *productServiceImpl) productReqToEntity(product *modelReq.Product) *entities.Product {
	return &entities.Product{
		ProductName:  product.ProductName,
		ProductPrice: product.ProductPrice,
	}

}

func (s *productServiceImpl) productEntityToRes(product *entities.Product) *modelRes.Product {
	return &modelRes.Product{
		ProductID:    product.ProductID,
		ProductName:  product.ProductName,
		ProductPrice: product.ProductPrice,
	}

}
