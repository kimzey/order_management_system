package entities

import (
	"time"
)

type Stock struct {
	StockID     uint64    `gorm:"primaryKey;autoIncrement;" json:"stockid"`
	ProductID   uint64    `gorm:"unique;not null;" json:"productid" validate:"required"`
	ProductName string    `gorm:"type:varchar(128);unique;not null;" json:"productname" validate:"required"`
	Quantity    uint      `gorm:"not null;" json:"quantity" validate:"required"`
	CreatedAt   time.Time `gorm:"not null;autoCreateTime;" `
	UpdatedAt   time.Time `gorm:"not null;autoUpdateTime;" `
}
