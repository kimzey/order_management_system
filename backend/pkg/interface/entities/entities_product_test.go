package entities

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProduct(t *testing.T) {
	product := Product{
		ProductID:    "1",
		ProductName:  "Product1",
		ProductPrice: 100,
	}

	assert.Equal(t, "1", product.ProductID)
	assert.Equal(t, "Product1", product.ProductName)
	assert.Equal(t, uint(100), product.ProductPrice)
}
