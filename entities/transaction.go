package entities

import (
	"time"
)

type Transaction struct {
	TransactionID uint64    `gorm:"primaryKey;autoIncrement;" json:"transactionid" validate:"required"`
	ProductName   string    `gorm:"type:varchar(128);not null;" json:"productname" validate:"required"`
	ProductPrice  uint      `gorm:"not null;" json:"productprice" validate:"required"`
	Quantity      uint      `gorm:"not null;" json:"quantity" validate:"required"`
	SumPrice      uint      `gorm:"not null;" json:"sumprice" validate:"required"`
	CreatedAt     time.Time `gorm:"not null;autoCreateTime;"`
	UpdatedAt     time.Time `gorm:"not null;autoUpdateTime;"`
}

//func (s *Transaction) calculatePrice(price uint, quantity uint) uint {
//	return price * quantity
//}
