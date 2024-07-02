package entities_test

import (
	"github.com/kizmey/order_management_system/pkg/interface/entities"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStock(t *testing.T) {
	stock := entities.Stock{
		ProductID: "1",
		Quantity:  10,
	}

	assert.Equal(t, "1", stock.ProductID)
	assert.Equal(t, uint(10), stock.Quantity)
}
