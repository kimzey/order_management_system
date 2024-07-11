package entities

type Transaction struct {
	TransactionID string
	//Quantity      uint
	SumPrice   uint
	IsDomestic bool
}

//const (
//	// Domestic price
//	Domestic = uint(100)
//
//	// NotDomestic price
//	NotDomestic = uint(500)
//)
//
//func (m *Transaction) CalculatePrice(price uint, quantity uint, isDomestic bool) uint {
//	if isDomestic {
//		return (price * quantity) + Domestic
//	} else {
//		return (price * quantity) + NotDomestic
//	}
//}
