package repository

import (
	"github.com/kizmey/order_management_system/database"
	"github.com/kizmey/order_management_system/entities"
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
	newProduct := new(entities.Product)

	if err := r.db.Connect().Create(product).Scan(newProduct).Error; err != nil {
		r.logger.Error("Creating item failed:", err.Error())
		return nil, err
	}
	return newProduct, nil
}

func (r *productRepositoryImpl) FindAll() (*[]entities.Product, error) {
	products := new([]entities.Product)

	if err := r.db.Connect().Find(products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (r *productRepositoryImpl) FindByID(id uint64) (*entities.Product, error) {
	product := new(entities.Product)

	if err := r.db.Connect().Where("product_id = ?", id).First(product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func (r *productRepositoryImpl) Update(id uint64, product *entities.Product) (*entities.Product, error) {
	newProduct := new(entities.Product)

	if err := r.db.Connect().Model(&entities.Product{}).Where(
		"product_id = ?", id,
	).Updates(
		product,
	).Scan(newProduct).Error; err != nil {
		r.logger.Error("Editing item failed:", err.Error())
		return nil, err
	}
	return newProduct, nil
}

func (r *productRepositoryImpl) Delete(id uint64) error {
	return r.db.Connect().Delete(&entities.Product{}, id).Error
}
