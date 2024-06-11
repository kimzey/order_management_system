package modelReq

type Stock struct {
	ProductID string `json:"productid" validate:"required"`
	Quantity  uint   `json:"quantity" validate:"required"`
}
