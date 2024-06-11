package main

import (
	"github.com/kizmey/order_management_system/config"
	"github.com/kizmey/order_management_system/database"
	model2 "github.com/kizmey/order_management_system/pkg/interface/model"
	"gorm.io/gorm"
)

func main() {
	conf := config.GettingConfig()
	db := database.NewPostgresDatabase(conf.Database)

	tx := db.Connect().Begin()

	productMigration(tx)
	stockMigration(tx)
	transactionMigration(tx)
	orderMigration(tx)

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		panic(err)
	}
}

func orderMigration(tx *gorm.DB) {
	err := tx.Migrator().CreateTable(&model2.Order{})
	if err != nil {
		return
	}
}

func productMigration(tx *gorm.DB) {
	err := tx.Migrator().CreateTable(&model2.Product{})
	if err != nil {
		return
	}
}

func stockMigration(tx *gorm.DB) {
	err := tx.Migrator().CreateTable(&model2.Stock{})
	if err != nil {
		return
	}
}

func transactionMigration(tx *gorm.DB) {
	err := tx.Migrator().CreateTable(&model2.Transaction{})
	if err != nil {
		return
	}
}
