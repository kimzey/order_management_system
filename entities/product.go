package entities

type Product struct {
	ProductID    uint64
	ProductName  string
	ProductPrice uint
}

//func (e *Product) ToProductModel() *model.Product {
//	return &model.Product{
//		Name:  e.ProductName,
//		Price: e.ProductPrice,
//	}
//}
