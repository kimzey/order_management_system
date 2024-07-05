package modelReq

type Order struct {
	TransactionID string `json:"transactionid" validate:"required"`
	Status        string `json:"status" `
}
