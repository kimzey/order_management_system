package entities

import (
	"github.com/kizmey/order_management_system/model"
)

type Order struct {
	OrderID       uint64 `json:"orderid" `
	TransactionID uint64 `json:"transactionid" `
	ProductID     uint64 `json:"productid" `
	IsDomestic    bool   `json:"isdomestic" `
	ProductName   string `json:"productname" `
	ProductPrice  uint   `json:"productprice" `
	Quantity      uint   `json:"quantity" `
	Status        string `json:"status"`
	SumPrice      uint   `json:"sumprice"`
}

func (e *Order) ToOrderModel() *model.Order {
	return &model.Order{
		TransactionID: e.TransactionID,
		ProductID:     e.ProductID,
		Status:        e.Status,
	}
}
