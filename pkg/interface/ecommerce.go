package _interface

import "github.com/kizmey/order_management_system/pkg/interface/entities"

type Ecommerce struct {
	Order               *entities.Order
	ProductID           string
	TransactionQuantity uint
}

func NewEcommerce(order *entities.Order, productid string, transactionQuantity uint) *Ecommerce {
	return &Ecommerce{
		Order:               order,
		ProductID:           productid,
		TransactionQuantity: transactionQuantity,
	}
}

type TransactionEcommerce struct {
	Tranasaction  *entities.Transaction
	Product       []entities.Product
	AddessProduct map[string]uint
}

func NewTransactionEcommerce(tranasaction *entities.Transaction, product []entities.Product, addessProduct map[string]uint) *TransactionEcommerce {
	return &TransactionEcommerce{
		Tranasaction:  tranasaction,
		Product:       product,
		AddessProduct: addessProduct,
	}
}
