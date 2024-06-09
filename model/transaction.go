package model

import (
	"github.com/kizmey/order_management_system/entities"
	"time"
)

type Transaction struct {
	ID         uint64    `gorm:"primaryKey;autoIncrement;" `
	ProductID  uint64    `gorm:"not null;" `
	Product    Product   `gorm:"foreignKey:ProductID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	IsDomestic bool      `gorm:"not null; default:false;"`
	Quantity   uint      `gorm:"not null;" `
	SumPrice   uint      `gorm:"not null;" `
	CreatedAt  time.Time `gorm:"not null;autoCreateTime;"`
	UpdatedAt  time.Time `gorm:"not null;autoUpdateTime;"`
}

func (m *Transaction) ToTransactionEntity() *entities.Transaction {
	return &entities.Transaction{
		TransactionID: m.ID,
		ProductID:     m.ProductID,
		ProductName:   m.Product.Name,
		ProductPrice:  m.Product.Price,
		Quantity:      m.Quantity,
		SumPrice:      m.SumPrice,
		IsDomestic:    m.IsDomestic,
	}
}

func ConvertModelsTransactionToEntities(transactions *[]Transaction) *[]entities.Transaction {
	entityTransaction := new([]entities.Transaction)

	for _, transaction := range *transactions {

		*entityTransaction = append(*entityTransaction, *transaction.ToTransactionEntity())
	}

	return entityTransaction
}
func (m *Transaction) CalculatePrice(price uint, quantity uint, isDomestic bool) uint {
	if isDomestic {
		return (price * quantity) + 40
	} else {
		return (price * quantity) + 200
	}
}
