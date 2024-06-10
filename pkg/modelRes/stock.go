package modelRes

type Stock struct {
	StockID   uint64 `json:"stockid"`
	ProductID uint64 `json:"productid" `
	Quantity  uint   `json:"quantity" `
}
