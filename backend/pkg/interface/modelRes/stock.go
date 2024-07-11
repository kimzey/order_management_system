package modelRes

type Stock struct {
	StockID   string `json:"stockid"`
	ProductID string `json:"productid" `
	Quantity  uint   `json:"quantity" `
}
