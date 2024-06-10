package modelReq

type Stock struct {
	ProductID uint64 `json:"productid" validate:"required"`
	Quantity  uint   `json:"quantity" validate:"required"`
}
