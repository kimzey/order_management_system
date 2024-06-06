package repository

import (
	"github.com/kizmey/order_management_system/database"
	"github.com/kizmey/order_management_system/entities"
	"github.com/kizmey/order_management_system/model"
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
	newStock := new(model.Stock)

	if err := r.db.Connect().Create(stock.ToStockModel()).Scan(newStock).Error; err != nil {
		r.logger.Error("Creating item failed:", err.Error())
		return nil, err
	}

	return newStock.ToStockEntity(), nil
}

func (r *stockRepositoryImpl) FindAll() (*[]entities.Stock, error) {
	stocks := new([]model.Stock)

	if err := r.db.Connect().Find(stocks).Error; err != nil {
		return nil, err
	}
	allStock := model.ConvertStockModelsToEntities(stocks)
	return allStock, nil
}

func (r *stockRepositoryImpl) CheckStockByProductId(productId uint64) (*entities.Stock, error) {
	stock := new(model.Stock)

	if err := r.db.Connect().Where("id = ?", productId).First(stock).Error; err != nil {
		return nil, err
	}
	return stock.ToStockEntity(), nil
}

func (r *stockRepositoryImpl) Update(stockid uint64, stock *entities.Stock) (*entities.Stock, error) {
	stocks := new(model.Stock)

	if err := r.db.Connect().Model(&model.Stock{}).Where(
		"id = ?", stockid,
	).Updates(
		stock,
	).Scan(stocks).Error; err != nil {
		r.logger.Error("Editing item failed:", err.Error())
		return nil, err
	}
	return stocks.ToStockEntity(), nil
}

func (r *stockRepositoryImpl) Delete(id uint64) error {
	return r.db.Connect().Delete(&model.Stock{}, id).Error
}
