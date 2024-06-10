package modelRes

type Transaction struct {
	ProductID   uint64 `json:"productid" `
	ProductName string `json:"productName"`
	Quantity    uint   `json:"quantity" `
	IsDomestic  bool   `json:"isdomestic" `
	SumPrice    uint   `json:"sumprice"`
}
