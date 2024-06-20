package product

import (
	"context"
	"github.com/kizmey/order_management_system/pkg/interface/entities"
	"github.com/kizmey/order_management_system/pkg/interface/modelReq"
	"github.com/kizmey/order_management_system/pkg/interface/modelRes"
	_ProductRepository "github.com/kizmey/order_management_system/pkg/repository/product"
)

type productServiceImpl struct {
	productRepository _ProductRepository.ProductRepository
}

func NewProductServiceImpl(productRepository _ProductRepository.ProductRepository) ProductService {
	return &productServiceImpl{productRepository: productRepository}
}

func (s *productServiceImpl) Create(ctx context.Context, product *modelReq.Product) (*modelRes.Product, error) {
	ctx, sp := tracer.Start(ctx, "productCreateService")
	defer sp.End()

	productEntity := s.productReqToEntity(product)

	productEntity, err := s.productRepository.Create(ctx, productEntity)
	if err != nil {
		return nil, err
	}
	return s.productEntityToRes(productEntity), nil
}

func (s *productServiceImpl) FindAll(ctx context.Context) (*[]modelRes.Product, error) {
	ctx, sp := tracer.Start(ctx, "productFindAllService")
	defer sp.End()

	products, err := s.productRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	productsRes := make([]modelRes.Product, 0)
	for _, product := range *products {
		productsRes = append(productsRes, *s.productEntityToRes(&product))
	}
	return &productsRes, nil
}

func (s *productServiceImpl) FindByID(ctx context.Context, id string) (*modelRes.Product, error) {
	ctx, sp := tracer.Start(ctx, "productFindByIdService")
	defer sp.End()

	product, err := s.productRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return s.productEntityToRes(product), nil
}
func (s *productServiceImpl) Update(ctx context.Context, id string, product *modelReq.Product) (*modelRes.Product, error) {
	ctx, sp := tracer.Start(ctx, "productUpdateService")
	defer sp.End()

	productEntity := s.productReqToEntity(product)
	productEntity, err := s.productRepository.Update(ctx, id, productEntity)
	if err != nil {
		return nil, err
	}
	return s.productEntityToRes(productEntity), nil
}

func (s *productServiceImpl) Delete(ctx context.Context, id string) (*modelRes.Product, error) {
	ctx, sp := tracer.Start(ctx, "")
	defer sp.End()

	product, err := s.productRepository.Delete(ctx, id)
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
