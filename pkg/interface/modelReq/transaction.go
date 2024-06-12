package modelReq

type Transaction struct {
	Product    map[string]uint `json:"product" validate:"required"`
	IsDomestic bool            `json:"isdomestic" `
}
