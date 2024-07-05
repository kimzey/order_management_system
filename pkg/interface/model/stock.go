package model

import (
	"github.com/kizmey/order_management_system/pkg/interface/entities"
	"time"
)

type Stock struct {
	ID        string    `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	ProductID string    `gorm:"unique;not null;" `
	Product   Product   `gorm:"foreignKey:ProductID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Quantity  uint      `gorm:"not null;" `
	CreatedAt time.Time `gorm:"not null;autoCreateTime;" `
	UpdatedAt time.Time `gorm:"not null;autoUpdateTime;" `
}

func (m *Stock) ToStockEntity() *entities.Stock {
	return &entities.Stock{
		StockID:   m.ID,
		ProductID: m.ProductID,
		Quantity:  m.Quantity,
	}
}

func ConvertStockModelsToEntities(stocks *[]Stock) *[]entities.Stock {
	entityStocks := new([]entities.Stock)

	for _, stock := range *stocks {
		*entityStocks = append(*entityStocks, *stock.ToStockEntity())
	}

	return entityStocks
}
