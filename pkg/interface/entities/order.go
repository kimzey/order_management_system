package entities

import (
	"errors"
)

type Order struct {
	OrderID       string
	TransactionID string
	Status        string
}

var (
	OrderStatus = []string{"New", "Paid", "Processing", "Done"}
)

func (m *Order) NextStatus() error {
	for i := 0; i < len(OrderStatus); i++ {

		if m.Status == OrderStatus[len(OrderStatus)-1] {
			return errors.New("order is already done")
		}

		if OrderStatus[i] == m.Status {
			m.Status = OrderStatus[i+1]
			return nil
		}
	}

	return errors.New("invalid order status")
}

func (m *Order) NextPaidToDone() error {

	if m.Status == "Paid" {
		m.Status = "Done"
		return nil
	}

	return errors.New("invalid order status")
}
