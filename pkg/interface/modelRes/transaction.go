package modelRes

type Transaction struct {
	TransactionID string `json:"transactionid" `
	ProductID     string `json:"productid" `
	ProductName   string `json:"productName"`
	Quantity      uint   `json:"quantity" `
	IsDomestic    bool   `json:"isdomestic" `
	SumPrice      uint   `json:"sumprice"`
}
