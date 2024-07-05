package modelRes

type Transaction struct {
	TransactionID string    `json:"transactionid" `
	IsDomestic    bool      `json:"isdomestic" `
	SumPrice      uint      `json:"sumprice"`
	Products      []Product `json:"products,omitempty"`
}
