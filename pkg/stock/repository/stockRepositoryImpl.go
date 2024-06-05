package repository

import (
	"github.com/kizmey/order_management_system/database"
	"github.com/kizmey/order_management_system/entities"
	"github.com/labstack/echo/v4"
)

type stockRepositoryImpl struct {
	db     database.Database
	logger echo.Logger
}

func NewStockRepositoryImpl(db database.Database, logger echo.Logger) StockRepository {
	return &stockRepositoryImpl{db: db, logger: logger}
}

func (r *stockRepositoryImpl) Create(stock *entities.Stock) (*entities.Stock, error) {
	newStock := new(entities.Stock)

	if err := r.db.Connect().Create(stock).Scan(newStock).Error; err != nil {
		r.logger.Error("Creating item failed:", err.Error())
		return nil, err
	}

	return newStock, nil
}

func (r *stockRepositoryImpl) FindAll() (*[]entities.Stock, error) {
	stocks := new([]entities.Stock)

	if err := r.db.Connect().Find(stocks).Error; err != nil {
		return nil, err
	}
	return stocks, nil
}

func (r *stockRepositoryImpl) CheckStockByProductId(productId uint64) (*entities.Stock, error) {
	stock := new(entities.Stock)

	if err := r.db.Connect().Where("product_id = ?", productId).First(stock).Error; err != nil {
		return nil, err
	}
	return stock, nil
}

func (r *stockRepositoryImpl) Update(stockid uint64, stock *entities.Stock) (*entities.Stock, error) {
	stocks := new(entities.Stock)

	if err := r.db.Connect().Model(&entities.Stock{}).Where(
		"stock_id = ?", stockid,
	).Updates(
		stock,
	).Scan(stocks).Error; err != nil {
		r.logger.Error("Editing item failed:", err.Error())
		return nil, err
	}
	return stocks, nil
}

func (r *stockRepositoryImpl) Delete(id uint64) error {
	return r.db.Connect().Delete(&entities.Stock{}, id).Error
}
