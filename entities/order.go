package entities

import (
	"errors"
	"time"
)

type Order struct {
	OrderID       uint64    `gorm:"primaryKey;autoIncrement;" json:"orderid" `
	TransactionID uint64    `gorm:"not null;" json:"transactionid" `
	IsDomestic    bool      `gorm:"not null; default:false;" json:"isdomestic" `
	ProductName   string    `gorm:"type:varchar(128);not null;" json:"productname" validate:"required"`
	ProductPrice  uint      `gorm:"not null;" json:"productprice" validate:"required"`
	Quantity      uint      `gorm:"not null;" json:"quantity" validate:"required"`
	Status        string    `gorm:"type:varchar(20);not null;default:New" json:"status"`
	CreatedAt     time.Time `gorm:"not null;autoCreateTime;"`
	UpdatedAt     time.Time `gorm:"not null;autoUpdateTime;"`
}

var (
	OrderStatus = []string{"New", "Paid", "Processing", "Done"}
)

func (e *Order) NextStatus() error {
	for i := 0; i < len(OrderStatus); i++ {

		if e.Status == OrderStatus[len(OrderStatus)-1] {
			return errors.New("order is already done")
		}

		if OrderStatus[i] == e.Status {
			e.Status = OrderStatus[i+1]
			return nil
		}
	}

	return errors.New("invalid order status")
}

func (e *Order) NextPaidToDone() error {

	if e.Status == "Paid" {
		e.Status = "Done"
		return nil
	}

	return errors.New("invalid order status")
}
