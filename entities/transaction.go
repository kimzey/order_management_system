package entities

type Transaction struct {
	TransactionID uint64
	ProductName   string
	ProductPrice  uint
	Quantity      uint
	SumPrice      uint
	IsDomestic    bool
}

//func (e *Transaction) ToTransactionModel() *model.Transaction {
//	return &model.Transaction{
//		ProductID:  e.ProductID,
//		Quantity:   e.Quantity,
//		SumPrice:   e.SumPrice,
//		IsDomestic: e.IsDomestic,
//	}
//}
