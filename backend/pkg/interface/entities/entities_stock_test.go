package entities

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStock(t *testing.T) {
	stock := Stock{
		ProductID: "1",
		Quantity:  10,
	}

	assert.Equal(t, "1", stock.ProductID)
	assert.Equal(t, uint(10), stock.Quantity)
}
