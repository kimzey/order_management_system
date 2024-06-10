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

//func ConvertStockModelsToEntities(stocks *[]Stock) *[]entities.Stock {
//	entityStocks := new([]entities.Stock)
//
//	for _, stock := range *stocks {
//		*entityStocks = append(*entityStocks, *stock.ToStockEntity())
//	}
//
//	return entityStocks
//}
