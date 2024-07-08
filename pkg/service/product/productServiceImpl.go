package product

import (
	"context"
	"errors"
	"github.com/kizmey/order_management_system/pkg/interface/entities"
	_ProductRepository "github.com/kizmey/order_management_system/pkg/repository/product"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type productServiceImpl struct {
	productRepository _ProductRepository.ProductRepository
}

func NewProductServiceImpl(productRepository _ProductRepository.ProductRepository) ProductService {
	return &productServiceImpl{productRepository: productRepository}
}

func (s *productServiceImpl) Create(ctx context.Context, product *entities.Product) (*entities.Product, error) {
	ctx, sp := tracer.Start(ctx, "productCreateService")
	defer sp.End()

	productEntity, err := s.productRepository.Create(ctx, product)
	if err != nil {
		return nil, err
	}

	s.SetProductSubAttributes(productEntity, sp)
	return productEntity, nil
}

func (s *productServiceImpl) FindAll(ctx context.Context) (*[]entities.Product, error) {
	ctx, sp := tracer.Start(ctx, "productFindAllService")
	defer sp.End()

	products, err := s.productRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	s.SetProductSubAttributes(products, sp)
	return products, nil
}

func (s *productServiceImpl) FindByID(ctx context.Context, id string) (*entities.Product, error) {
	ctx, sp := tracer.Start(ctx, "productFindByIdService")
	defer sp.End()

	product, err := s.productRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	s.SetProductSubAttributes(product, sp)
	return product, nil
}
func (s *productServiceImpl) Update(ctx context.Context, id string, product *entities.Product) (*entities.Product, error) {
	ctx, sp := tracer.Start(ctx, "productUpdateService")
	defer sp.End()

	productEntity, err := s.productRepository.Update(ctx, id, product)
	if err != nil {
		return nil, err
	}

	s.SetProductSubAttributes(productEntity, sp)
	return productEntity, nil
}

func (s *productServiceImpl) Delete(ctx context.Context, id string) (*entities.Product, error) {
	ctx, sp := tracer.Start(ctx, "")
	defer sp.End()

	product, err := s.productRepository.Delete(ctx, id)
	if err != nil {
		return nil, err
	}

	s.SetProductSubAttributes(product, sp)
	return product, nil
}

func (s *productServiceImpl) SetProductSubAttributes(productData any, sp trace.Span) {
	if products, ok := productData.(*[]entities.Product); ok {
		var productIDs []string
		var productNames []string
		var productPrices []int

		for _, product := range *products {
			productIDs = append(productIDs, product.ProductID)
			productNames = append(productNames, product.ProductName)
			productPrices = append(productPrices, int(product.ProductPrice))
		}

		sp.SetAttributes(
			attribute.StringSlice("ProductID", productIDs),
			attribute.StringSlice("ProductName", productNames),
			attribute.IntSlice("ProductPrice", productPrices),
		)
	} else if product, ok := productData.(*entities.Product); ok {
		sp.SetAttributes(
			attribute.String("ProductID", product.ProductID),
			attribute.String("ProductName", product.ProductName),
			attribute.Int("ProductPrice", int(product.ProductPrice)),
		)
	} else {
		sp.RecordError(errors.New("invalid type"))
	}
}
