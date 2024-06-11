package modelReq

type Order struct {
	TransactionID uint64 `json:"transactionid" validate:"required"`
	Status        string `json:"status" `
}
