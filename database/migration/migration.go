package main

import (
	"github.com/kizmey/order_management_system/config"
	"github.com/kizmey/order_management_system/database"
	"github.com/kizmey/order_management_system/entities"
	"gorm.io/gorm"
)

func main() {
	conf := config.GettingConfig()
	db := database.NewPostgresDatabase(conf.Database)

	tx := db.Connect().Begin()

	orderMigration(tx)
	productMigration(tx)
	stockMigration(tx)
	transactionMigration(tx)

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		panic(err)
	}
}

func orderMigration(tx *gorm.DB) {
	err := tx.Migrator().CreateTable(&entities.Order{})
	if err != nil {
		return
	}
}

func productMigration(tx *gorm.DB) {
	err := tx.Migrator().CreateTable(&entities.Product{})
	if err != nil {
		return
	}
}

func stockMigration(tx *gorm.DB) {
	err := tx.Migrator().CreateTable(&entities.Stock{})
	if err != nil {
		return
	}
}

func transactionMigration(tx *gorm.DB) {
	err := tx.Migrator().CreateTable(&entities.Transaction{})
	if err != nil {
		return
	}
}
