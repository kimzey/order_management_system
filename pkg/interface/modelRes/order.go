package modelRes

type Order struct {
	OrderID       string `json:"orderid"  `
	TransactionID string `json:"transactionid" `
	Status        string `json:"status"`
}
