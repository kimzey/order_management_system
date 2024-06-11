package repository

import (
	"github.com/kizmey/order_management_system/database"
	"github.com/kizmey/order_management_system/entities"
	"github.com/kizmey/order_management_system/model"
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
