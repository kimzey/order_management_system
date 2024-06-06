package model

import (
	"errors"
	"github.com/kizmey/order_management_system/entities"
	"time"
)

type Order struct {
	ID            uint64      `gorm:"primaryKey;autoIncrement;"`
	TransactionID uint64      `gorm:"not null;" `
	Transaction   Transaction `gorm:"foreignKey:TransactionID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ProductID     uint64      `gorm:"not null;" `
	Product       Product     `gorm:"foreignKey:ProductID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Status        string      `gorm:"type:varchar(20);not null;default:New"`
	CreatedAt     time.Time   `gorm:"not null;autoCreateTime;"`
	UpdatedAt     time.Time   `gorm:"not null;autoUpdateTime;"`
}

func (m *Order) ToOrderEntity() *entities.Order {
	return &entities.Order{
		OrderID:       m.ID,
		TransactionID: m.TransactionID,
		IsDomestic:    m.Transaction.IsDomestic,
		ProductName:   m.Product.Name,
		ProductPrice:  m.Product.Price,
		Quantity:      m.Transaction.Quantity,
		Status:        m.Status,
		SumPrice:      m.Transaction.SumPrice,
	}
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
