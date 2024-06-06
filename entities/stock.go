package entities

import "github.com/kizmey/order_management_system/model"

type Stock struct {
	StockID   uint64 `json:"stockid"`
	ProductID uint64 `json:"productid" `
	Quantity  uint   `json:"quantity" `
}

func (e *Stock) ToStockModel() *model.Stock {
	return &model.Stock{
		ProductID: e.ProductID,
		Quantity:  e.Quantity,
	}
}
