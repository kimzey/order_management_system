package _interface

import (
	"github.com/kizmey/order_management_system/pkg/interface/entities"
)

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

type TransactionEcommerce struct {
	Tranasaction  *entities.Transaction
	Product       []entities.Product
	AddessProduct map[string]uint
}

const (
	// Domestic price
	Domestic = uint(100)
	// NotDomestic price
	NotDomestic = uint(500)
)

func (m *TransactionEcommerce) CalculatePrice() uint {
	m.Tranasaction.SumPrice = 0
	for _, product := range m.Product {
		m.Tranasaction.SumPrice += (product.ProductPrice * m.AddessProduct[product.ProductID])
	}

	if m.Tranasaction.IsDomestic {
		m.Tranasaction.SumPrice += Domestic
	} else {
		m.Tranasaction.SumPrice += NotDomestic
	}

	return m.Tranasaction.SumPrice
}

func NewTransactionEcommerce(tranasaction *entities.Transaction, product []entities.Product, addessProduct map[string]uint) *TransactionEcommerce {
	return &TransactionEcommerce{
		Tranasaction:  tranasaction,
		Product:       product,
		AddessProduct: addessProduct,
	}
}
