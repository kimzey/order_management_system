package entities

import "github.com/kizmey/order_management_system/model"

type Product struct {
	ProductID    uint64 `json:"ProductID" `
	ProductName  string `json:"ProductName"`
	ProductPrice uint   `json:"ProductPrice" `
}

func (e *Product) ToProductModel() *model.Product {
	return &model.Product{
		Name:  e.ProductName,
		Price: e.ProductPrice,
	}
}
