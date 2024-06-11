package modelReq

type Transaction struct {
	ProductID  string `json:"productid" validate:"required"`
	Quantity   uint   `json:"quantity" validate:"required"`
	IsDomestic bool   `json:"isdomestic" validate:"required"`
}
