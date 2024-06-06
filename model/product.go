package model

import (
	"github.com/kizmey/order_management_system/entities"
	"time"
)

type Product struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement;" `
	Name      string    `gorm:"type:varchar(128);not null;" `
	Price     uint      `gorm:"not null;"`
	CreatedAt time.Time `gorm:"not null;autoCreateTime;" `
	UpdatedAt time.Time `gorm:"not null;autoUpdateTime;" `
}

func (m *Product) ToProductEntity() *entities.Product {
	return &entities.Product{
		ProductID:    m.ID,
		ProductName:  m.Name,
		ProductPrice: m.Price,
	}
}
func ConvertProductModelsToEntities(products *[]Product) *[]entities.Product {
	entityProducts := new([]entities.Product)

	for _, product := range *products {
		*entityProducts = append(*entityProducts, *product.ToProductEntity())
	}

	return entityProducts
}
