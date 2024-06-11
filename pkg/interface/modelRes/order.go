package modelRes

type Order struct {
	OrderID       uint64 `json:"orderid"  `
	TransactionID uint64 `json:"transactionid" `
	ProductID     uint64 `json:"productid" `
	Status        string `json:"status"`
}
