package entities

import (
	"time"
)

type Product struct {
	ProductID    uint64    `gorm:"primaryKey;autoIncrement;" json:"ProductID" `
	ProductName  string    `gorm:"type:varchar(128);not null;" json:"ProductName" validate:"required"`
	ProductPrice uint      `gorm:"not null;" json:"ProductPrice" validate:"required"`
	CreatedAt    time.Time `gorm:"not null;autoCreateTime;" `
	UpdatedAt    time.Time `gorm:"not null;autoUpdateTime;" `
}
