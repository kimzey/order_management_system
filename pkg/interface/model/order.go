package model

import (
	"github.com/kizmey/order_management_system/pkg/interface/entities"
	"time"
)

type Order struct {
	ID            uint64      `gorm:"primaryKey;autoIncrement;"`
	TransactionID uint64      `gorm:"not null; unique;" `
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
		Status:        m.Status,
	}
}
