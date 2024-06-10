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

//func ConvertProductModelsToEntities(products *[]Product) *[]entities.Product {
//	entityProducts := new([]entities.Product)
//
//	for _, product := range *products {
//		*entityProducts = append(*entityProducts, *product.ToProductEntity())
//	}
//
//	return entityProducts
//}
