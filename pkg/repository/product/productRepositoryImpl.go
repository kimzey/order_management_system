package product

import (
	"context"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"

	"errors"
	"fmt"
	"github.com/kizmey/order_management_system/database"
	"github.com/kizmey/order_management_system/pkg/interface/entities"
	"github.com/kizmey/order_management_system/pkg/interface/model"
)

type productRepositoryImpl struct {
	db database.Database
}

func NewProductRepositoryImpl(db database.Database) ProductRepository {
	return &productRepositoryImpl{db: db}

}

func (r *productRepositoryImpl) Create(ctx context.Context, product *entities.Product) (*entities.Product, error) {
	ctx, sp := tracer.Start(ctx, "productCreateRepository")
	defer sp.End()

	modelProduct := r.ToProductModel(product)
	newProduct := new(model.Product)

	if err := r.db.Connect().Create(modelProduct).Scan(newProduct).Error; err != nil {
		return nil, errors.New(fmt.Sprintf("failed to create product"))
	}

	productEntity := newProduct.ToProductEntity()
	r.SetProductSubAttributes(newProduct, sp)
	return productEntity, nil
}

func (r *productRepositoryImpl) FindAll(ctx context.Context) (*[]entities.Product, error) {
	ctx, sp := tracer.Start(ctx, "productFindByIdRepository")
	defer sp.End()

	products := new([]model.Product)

	if err := r.db.Connect().Find(products).Error; err != nil {
		return nil, errors.New(fmt.Sprintf("failed to find all products"))
	}
	allProduct := model.ConvertProductModelsToEntities(products)
	r.SetProductSubAttributes(allProduct, sp)
	return allProduct, nil
}

func (r *productRepositoryImpl) FindByID(ctx context.Context, id string) (*entities.Product, error) {
	ctx, sp := tracer.Start(ctx, "productFindByIdRepository")
	defer sp.End()

	product := new(model.Product)

	if err := r.db.Connect().Where("id = ?", id).First(product).Error; err != nil {
		return nil, errors.New(fmt.Sprintf("failed to find product"))
	}

	productEntity := product.ToProductEntity()
	r.SetProductSubAttributes(product, sp)
	return productEntity, nil
}

func (r *productRepositoryImpl) Update(ctx context.Context, id string, product *entities.Product) (*entities.Product, error) {
	ctx, sp := tracer.Start(ctx, "productFindByIdRepository")
	defer sp.End()

	newProduct := new(model.Product)
	productModel := r.ToProductModel(product)

	if err := r.db.Connect().Model(&productModel).Where(
		"id = ?", id,
	).Updates(
		productModel,
	).Scan(newProduct).Error; err != nil {
		return nil, errors.New(fmt.Sprintf("failed to update product"))
	}

	productEntity := newProduct.ToProductEntity()
	r.SetProductSubAttributes(product, sp)
	return productEntity, nil
}

func (r *productRepositoryImpl) Delete(ctx context.Context, id string) (*entities.Product, error) {
	ctx, sp := tracer.Start(ctx, "productDeleteRepository")
	defer sp.End()

	product := new(model.Product)
	if err := r.db.Connect().Where("id = ?", id).First(&product).Delete(&product).Error; err != nil {
		return nil, errors.New(fmt.Sprintf("failed to delete product"))
	}

	productEntity := product.ToProductEntity()
	r.SetProductSubAttributes(product, sp)
	return productEntity, nil
}

func (r *productRepositoryImpl) ToProductModel(e *entities.Product) *model.Product {
	return &model.Product{
		Name:  e.ProductName,
		Price: e.ProductPrice,
	}
}

func (r *productRepositoryImpl) SetProductSubAttributes(productData any, sp trace.Span) {
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
