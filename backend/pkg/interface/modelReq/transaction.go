package modelReq

type ProductItem struct {
	ProductID string `json:"product_id" validate:"required"`
	Quantity  uint   `json:"quantity" validate:"required,min=1"`
}

type Transaction struct {
	Product    []ProductItem `json:"product" validate:"required,dive"`
	IsDomestic bool          `json:"isdomestic"`
}
