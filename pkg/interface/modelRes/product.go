package modelRes

type Product struct {
	ProductID    string `json:"productID" `
	ProductName  string `json:"productName"`
	ProductPrice uint   `json:"productPrice,omitempty" `
	Quantity     uint   `json:"quantity,omitempty"`
}
