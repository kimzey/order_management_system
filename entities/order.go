package entities

type Order struct {
	OrderID      uint64
	ProductName  string
	ProductPrice uint
	IsDomestic   bool
	Quantity     uint
	Status       string
	SumPrice     uint
}

//func (e *Order) ToOrderModel() *model.Order {
//	return &model.Order{
//		TransactionID: e.TransactionID,
//		ProductID:     e.ProductID,
//		Status:        e.Status,
//	}
//}
