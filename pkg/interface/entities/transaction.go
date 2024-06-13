package entities

type Transaction struct {
	TransactionID string
	Quantity      uint
	SumPrice      uint
	IsDomestic    bool
	Address       Address
}
type Address struct {
	StreetAddress string `validate:"required,min=4,max=128"`
	SubDistrict   string `validate:"required,min=2,max=64"`
	District      string `validate:"required,min=2,max=64"`
	Province      string `validate:"required,min=2,max=64"`
	PostalCode    string `validate:"required,min=3,max=8"`
	Country       string `validate:"required,min=3,max=64"`
}

func (m *Transaction) CalculatePrice(price uint, quantity uint, isDomestic bool) uint {
	if isDomestic {
		return (price * quantity) + 40
	} else {
		return (price * quantity) + 200
	}
}
