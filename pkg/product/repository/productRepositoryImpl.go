package repository

import (
	"github.com/kizmey/order_management_system/database"
	"github.com/kizmey/order_management_system/entities"
	"github.com/kizmey/order_management_system/model"
	"github.com/labstack/echo/v4"
)

type productRepositoryImpl struct {
	db     database.Database
	logger echo.Logger
}

func NewProductRepositoryImpl(db database.Database, logger echo.Logger) ProductRepository {
	return &productRepositoryImpl{db: db, logger: logger}

}

func (r *productRepositoryImpl) Create(product *entities.Product) (*entities.Product, error) {
	modelProduct := ToProductModel(product)
	newProduct := new(model.Product)

	if err := r.db.Connect().Create(modelProduct).Scan(newProduct).Error; err != nil {
		r.logger.Error("Creating item failed:", err.Error())
		return nil, err
	}
	return newProduct.ToProductEntity(), nil
}

func (r *productRepositoryImpl) FindAll() (*[]entities.Product, error) {
	products := new([]model.Product)

	if err := r.db.Connect().Find(products).Error; err != nil {
		return nil, err
	}
	allProduct := model.ConvertProductModelsToEntities(products)
	return allProduct, nil
}

func (r *productRepositoryImpl) FindByID(id uint64) (*entities.Product, error) {
	product := new(model.Product)

	if err := r.db.Connect().Where("id = ?", id).First(product).Error; err != nil {
		r.logger.Error("Failed to find product:", err.Error())
		return nil, err
	}
	return product.ToProductEntity(), nil
}

func (r *productRepositoryImpl) Update(id uint64, product *entities.Product) (*entities.Product, error) {
	newProduct := new(model.Product)
	productModel := ToProductModel(product)

	if err := r.db.Connect().Model(&productModel).Where(
		"id = ?", id,
	).Updates(
		productModel,
	).Scan(newProduct).Error; err != nil {
		r.logger.Error("Editing item failed:", err.Error())
		return nil, err
	}
	return newProduct.ToProductEntity(), nil
}

func (r *productRepositoryImpl) Delete(id uint64) error {
	return r.db.Connect().Delete(&model.Product{}, id).Error
}

func ToProductModel(e *entities.Product) *model.Product {
	return &model.Product{
		Name:  e.ProductName,
		Price: e.ProductPrice,
	}
}
