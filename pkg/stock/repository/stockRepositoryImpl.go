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
	modelStock := ToStockModel(stock)
	newStock := new(model.Stock)

	if err := r.db.Connect().Create(modelStock).Scan(newStock).Error; err != nil {
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
	//fmt.Println("productId: ", productId)
	if err := r.db.Connect().Preload("Product").Where("product_id = ?", productId).First(stock).Error; err != nil {
		return nil, err
	}
	//fmt.Println("stock: ", stock)
	return stock.ToStockEntity(), nil
}

func (r *stockRepositoryImpl) Update(stockid uint64, stock *entities.Stock) (*entities.Stock, error) {
	stocks := new(model.Stock)
	modelStock := ToStockModel(stock)

	if err := r.db.Connect().Model(&modelStock).Where(
		"id = ?", stockid,
	).Updates(
		modelStock,
	).Scan(stocks).Error; err != nil {
		r.logger.Error("Editing item failed:", err.Error())
		return nil, err
	}
	//fmt.Println("stocks: ", stocks)
	return stocks.ToStockEntity(), nil
}

func (r *stockRepositoryImpl) Delete(id uint64) error {
	return r.db.Connect().Delete(&model.Stock{}, id).Error
}

func ToStockModel(e *entities.Stock) *model.Stock {
	return &model.Stock{
		ProductID: e.ProductID,
		Quantity:  e.Quantity,
	}
}
