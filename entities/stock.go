package entities

type Stock struct {
	StockID   uint64
	ProductID uint64
	Quantity  uint
}

//func (e *Stock) ToStockModel() *model.Stock {
//	return &model.Stock{
//		ProductID: e.ProductID,
//		Quantity:  e.Quantity,
//	}
//}
