package entities_test

import (
	"github.com/kizmey/order_management_system/pkg/interface/aggregation"
	"testing"

	"github.com/kizmey/order_management_system/pkg/interface/entities"
	"github.com/stretchr/testify/assert"
)

func TestTransactionecommerceCalculateprice(t *testing.T) {
	tests := []struct {
		name          string
		isDomestic    bool
		products      []entities.Product
		addessProduct map[string]uint
		expectedPrice uint
	}{
		{
			name:       "domestic order",
			isDomestic: true,
			products: []entities.Product{
				{ProductID: "1", ProductName: "Product1", ProductPrice: 100},
				{ProductID: "2", ProductName: "Product2", ProductPrice: 200},
			},
			addessProduct: map[string]uint{
				"1": 1,
				"2": 2,
			},
			expectedPrice: (100*1 + 200*2) + aggregation.Domestic,
		},
		{
			name:       "non-domestic order",
			isDomestic: false,
			products: []entities.Product{
				{ProductID: "1", ProductName: "Product1", ProductPrice: 100},
				{ProductID: "2", ProductName: "Product2", ProductPrice: 200},
			},
			addessProduct: map[string]uint{
				"1": 1,
				"2": 2,
			},
			expectedPrice: (100*1 + 200*2) + aggregation.NotDomestic,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transaction := &entities.Transaction{IsDomestic: tt.isDomestic}
			transactionEcommerce := aggregation.NewTransactionEcommerce(transaction, tt.products, tt.addessProduct)

			calculatedPrice := transactionEcommerce.CalculatePrice()
			assert.Equal(t, tt.expectedPrice, calculatedPrice)
		})
	}
}
