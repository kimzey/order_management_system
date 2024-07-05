package entities_test

import (
	"github.com/kizmey/order_management_system/pkg/interface/entities"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProduct(t *testing.T) {
	product := entities.Product{
		ProductID:    "1",
		ProductName:  "Product1",
		ProductPrice: 100,
	}

	assert.Equal(t, "1", product.ProductID)
	assert.Equal(t, "Product1", product.ProductName)
	assert.Equal(t, uint(100), product.ProductPrice)
}
