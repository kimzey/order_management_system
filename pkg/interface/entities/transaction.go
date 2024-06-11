package entities

type Transaction struct {
	TransactionID string
	ProductID     string
	ProductName   string
	ProductPrice  uint
	Quantity      uint
	SumPrice      uint
	IsDomestic    bool
}

func (m *Transaction) CalculatePrice(price uint, quantity uint, isDomestic bool) uint {
	if isDomestic {
		return (price * quantity) + 40
	} else {
		return (price * quantity) + 200
	}
}
