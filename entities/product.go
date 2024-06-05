package entities

import (
	"time"
)

type Product struct {
	ProductID    uint64    `gorm:"primaryKey;autoIncrement;" json:"ProductID" validate:"required"`
	ProductName  string    `gorm:"type:varchar(128);not null;" json:"ProductName" validate:"required"`
	ProductPrice uint      `gorm:"not null;" json:"ProductPrice" validate:"required"`
	CreatedAt    time.Time `gorm:"not null;autoCreateTime;" json:"CreatedAt" validate:"required"`
	UpdatedAt    time.Time `gorm:"not null;autoUpdateTime;" json:"UpdatedAt" validate:"required"`
}
