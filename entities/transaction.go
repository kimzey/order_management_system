package entities

import "github.com/kizmey/order_management_system/model"

type Transaction struct {
	TransactionID uint64 `json:"transactionid" `
	ProductID     uint64 `json:"ProductID" `
	ProductName   string `json:"productname" `
	ProductPrice  uint   `json:"productprice" `
	Quantity      uint   `json:"quantity" `
	SumPrice      uint   `json:"sumprice" `
	IsDomestic    bool   `json:"isdomestic" `
}

func (e *Transaction) ToTransactionModel() *model.Transaction {
	return &model.Transaction{
		ProductID:  e.ProductID,
		Quantity:   e.Quantity,
		SumPrice:   e.SumPrice,
		IsDomestic: e.IsDomestic,
	}
}
