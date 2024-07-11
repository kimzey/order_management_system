package entities_test

import (
	"github.com/kizmey/order_management_system/pkg/interface/entities"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTransaction(t *testing.T) {
	transaction := &entities.Transaction{
		TransactionID: "1",
		SumPrice:      100,
		IsDomestic:    true,
	}
	assert.Equal(t, transaction.TransactionID, "1")
	assert.Equal(t, transaction.SumPrice, uint(100))
	assert.Equal(t, transaction.IsDomestic, true)
}
