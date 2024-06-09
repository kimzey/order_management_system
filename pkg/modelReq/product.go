package entities

type Product struct {
	ProductID    uint64 `json:"productID" `
	ProductName  string `json:"productName"`
	ProductPrice uint   `json:"productPrice" `
}

//func (e *Product) ToProductModel() *model.Product {
//	return &model.Product{
//		Name:  e.ProductName,
//		Price: e.ProductPrice,
//	}
//}
