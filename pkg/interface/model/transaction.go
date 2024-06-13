package model

import (
	"github.com/kizmey/order_management_system/pkg/interface/entities"
	"time"
)

type Transaction struct {
	ID         string       `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	IsDomestic bool         `gorm:"not null; default:false;"`
	Address    AddressModel `gorm:"embedded"`
	SumPrice   uint         `gorm:"not null;" `
	ProductID  string       `gorm:"not null;" `
	Products   []Product    `gorm:"many2many:transaction_products;"`
	CreatedAt  time.Time    `gorm:"not null;autoCreateTime;"`
	UpdatedAt  time.Time    `gorm:"not null;autoUpdateTime;"`
}

type AddressModel struct {
	StreetAddress string `gorm:"column:street_address"`
	SubDistrict   string `gorm:"column:sub_district"`
	District      string `gorm:"column:district"`
	Province      string `gorm:"column:province"`
	PostalCode    string `gorm:"column:postal_code"`
	Country       string `gorm:"column:country"`
}

type TransactionProduct struct {
	ID            string      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	TransactionID string      `gorm:"not null;" `
	Transaction   Transaction `gorm:"foreignKey:TransactionID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ProductID     string      `gorm:"not null;" `
	Product       Product     `gorm:"foreignKey:ProductID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Quantity      uint        `gorm:" not null; default:1" `
	//CreatedAt     time.Time   `gorm:"not null;autoCreateTime;"`
	//UpdatedAt     time.Time   `gorm:"not null;autoUpdateTime;"`
}

func (m *Transaction) ToTransactionEntity() *entities.Transaction {
	return &entities.Transaction{
		TransactionID: m.ID,
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
