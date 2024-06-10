package modelReq

type Product struct {
	ProductName  string `json:"productName" validate:"required"`
	ProductPrice uint   `json:"productPrice" validate:"required"`
}
