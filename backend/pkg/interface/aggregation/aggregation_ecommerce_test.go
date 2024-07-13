package aggregation

import (
	"testing"

	"github.com/kizmey/order_management_system/pkg/interface/entities"
	"github.com/stretchr/testify/assert"
)

func TestNewEcommerce(t *testing.T) {
	order := &entities.Order{OrderID: "1", Status: "New"}
	products := []entities.Product{
		{ProductID: "1", ProductName: "Product1", ProductPrice: 100},
		{ProductID: "2", ProductName: "Product2", ProductPrice: 200},
	}
	quantities := []uint{1, 2}

	ecommerce := NewEcommerce(order, products, quantities)

	assert.Equal(t, order, ecommerce.Order)
	assert.Equal(t, products, ecommerce.Product)
	assert.Equal(t, quantities, ecommerce.Quantity)
}
