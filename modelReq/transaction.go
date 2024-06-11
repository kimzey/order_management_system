package modelReq

type Transaction struct {
	ProductID  uint64 `json:"productid" validate:"required"`
	Quantity   uint   `json:"quantity" validate:"required"`
	IsDomestic bool   `json:"isdomestic" validate:"required"`
}
