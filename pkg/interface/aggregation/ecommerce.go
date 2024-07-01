package aggregation

import "github.com/kizmey/order_management_system/pkg/interface/entities"

type Ecommerce struct {
	Order    *entities.Order
	Product  []entities.Product
	Quantity []uint
}

func NewEcommerce(order *entities.Order, products []entities.Product, quantity []uint) *Ecommerce {
	return &Ecommerce{
		Order:    order,
		Product:  products,
		Quantity: quantity,
	}
}
