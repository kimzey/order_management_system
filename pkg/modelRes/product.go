package modelRes

type Product struct {
	ProductID    uint64 `json:"productID" `
	ProductName  string `json:"productName"`
	ProductPrice uint   `json:"productPrice" `
}
