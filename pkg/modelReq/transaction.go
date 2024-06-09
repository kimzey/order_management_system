package entities

type Transaction struct {
	TransactionID uint64 `json:"transactionid" `
	ProductID     uint64 `json:"productid" `
	ProductName   string `json:"productname" `
	ProductPrice  uint   `json:"productprice" `
	Quantity      uint   `json:"quantity" `
	SumPrice      uint   `json:"sumprice" `
	IsDomestic    bool   `json:"isdomestic" `
}

//func (e *Transaction) ToTransactionModel() *model.Transaction {
//	return &model.Transaction{
//		ProductID:  e.ProductID,
//		Quantity:   e.Quantity,
//		SumPrice:   e.SumPrice,
//		IsDomestic: e.IsDomestic,
//	}
//}
