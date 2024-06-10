package modelReq

type Order struct {
	TransactionID uint64 `json:"transactionid" validate:"required"`
	ProductID     uint64 `json:"productid" validate:"required"`
	Status        string `json:"status" validate:"required"`
}
