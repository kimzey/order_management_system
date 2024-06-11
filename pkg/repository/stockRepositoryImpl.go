package repository

import (
	"errors"
	"fmt"
	"github.com/kizmey/order_management_system/database"
	"github.com/kizmey/order_management_system/pkg/interface/entities"
	"github.com/kizmey/order_management_system/pkg/interface/model"
)

type stockRepositoryImpl struct {
	db database.Database
}

func NewStockRepositoryImpl(db database.Database) StockRepository {
	return &stockRepositoryImpl{db: db}
}

func (r *stockRepositoryImpl) Create(stock *entities.Stock) (*entities.Stock, error) {
	modelStock := ToStockModel(stock)
	newStock := new(model.Stock)

	if err := r.db.Connect().Create(modelStock).Scan(newStock).Error; err != nil {
		return nil, errors.New(fmt.Sprintf("failed to create stock: %s", err.Error()))
	}

	return newStock.ToStockEntity(), nil
}

func (r *stockRepositoryImpl) FindAll() (*[]entities.Stock, error) {
	stocks := new([]model.Stock)

	if err := r.db.Connect().Find(stocks).Error; err != nil {
		return nil, errors.New(fmt.Sprintf("failed to find all stock: %s", err.Error()))
	}
	allStock := model.ConvertStockModelsToEntities(stocks)
	return allStock, nil
}

func (r *stockRepositoryImpl) CheckStockByProductId(productId uint64) (*entities.Stock, error) {
	stock := new(model.Stock)
	//fmt.Println("productId: ", productId)
	if err := r.db.Connect().Preload("Product").Where("product_id = ?", productId).First(stock).Error; err != nil {
		return nil, errors.New(fmt.Sprintf("failed to find stock: %s", err.Error()))
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
		return nil, errors.New(fmt.Sprintf("failed to update stock: %s", err.Error()))
	}
	//fmt.Println("stocks: ", stocks)
	return stocks.ToStockEntity(), nil
}

func (r *stockRepositoryImpl) Delete(id uint64) error {
	if err := r.db.Connect().Delete(&model.Stock{}, id).Error; err != nil {
		return errors.New(fmt.Sprintf("failed to delete stock: %s", err.Error()))
	}
	return nil
}

func ToStockModel(e *entities.Stock) *model.Stock {
	return &model.Stock{
		ProductID: e.ProductID,
		Quantity:  e.Quantity,
	}
}
