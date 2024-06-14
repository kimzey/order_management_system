package repository

import (
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

func (r *productRepositoryImpl) Create(product *entities.Product) (*entities.Product, error) {
	modelProduct := ToProductModel(product)
	newProduct := new(model.Product)

	if err := r.db.Connect().Create(modelProduct).Scan(newProduct).Error; err != nil {
		return nil, errors.New(fmt.Sprintf("failed to create product"))
	}
	return newProduct.ToProductEntity(), nil
}

func (r *productRepositoryImpl) FindAll() (*[]entities.Product, error) {
	products := new([]model.Product)

	if err := r.db.Connect().Find(products).Error; err != nil {
		return nil, errors.New(fmt.Sprintf("failed to find all products"))
	}
	allProduct := model.ConvertProductModelsToEntities(products)
	return allProduct, nil
}

func (r *productRepositoryImpl) FindByID(id string) (*entities.Product, error) {
	product := new(model.Product)

	if err := r.db.Connect().Where("id = ?", id).First(product).Error; err != nil {
		return nil, errors.New(fmt.Sprintf("failed to find product",))
	}
	return product.ToProductEntity(), nil
}

func (r *productRepositoryImpl) Update(id string, product *entities.Product) (*entities.Product, error) {
	newProduct := new(model.Product)
	productModel := ToProductModel(product)

	if err := r.db.Connect().Model(&productModel).Where(
		"id = ?", id,
	).Updates(
		productModel,
	).Scan(newProduct).Error; err != nil {
		return nil, errors.New(fmt.Sprintf("failed to update product"))
	}
	return newProduct.ToProductEntity(), nil
}

func (r *productRepositoryImpl) Delete(id string) error {

	if err := r.db.Connect().Where("id = ?", id).Delete(&model.Product{}).Error; err != nil {
		return errors.New(fmt.Sprintf("failed to delete product"))
	}
	return nil
}

func ToProductModel(e *entities.Product) *model.Product {
	return &model.Product{
		Name:  e.ProductName,
		Price: e.ProductPrice,
	}
}
