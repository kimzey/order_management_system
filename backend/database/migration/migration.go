package migration

import (
	"github.com/kizmey/order_management_system/config"
	"github.com/kizmey/order_management_system/database"
	"github.com/kizmey/order_management_system/pkg/interface/model"
	"gorm.io/gorm"
)

func GettingMigration() {
	conf := config.GettingConfig()
	db := database.NewPostgresDatabase(conf.Database)

	tx := db.Connect().Begin()

	//if err := tx.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";").Error; err != nil {
	//	tx.Rollback()
	//	panic(err)
	//}

	productMigration(tx)
	stockMigration(tx)
	transactionMigration(tx)
	transactionproductMigration(tx)
	orderMigration(tx)

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		panic(err)
	}
}

func orderMigration(tx *gorm.DB) {
	err := tx.Migrator().CreateTable(&model.Order{})
	if err != nil {
		return
	}
}

func productMigration(tx *gorm.DB) {
	err := tx.Migrator().CreateTable(&model.Product{})
	if err != nil {
		return
	}
}

func stockMigration(tx *gorm.DB) {
	err := tx.Migrator().CreateTable(&model.Stock{})
	if err != nil {
		return
	}
}

func transactionMigration(tx *gorm.DB) {
	err := tx.Migrator().CreateTable(&model.Transaction{})
	if err != nil {
		return
	}
}
func transactionproductMigration(tx *gorm.DB) {
	err := tx.Migrator().CreateTable(&model.TransactionProduct{})
	if err != nil {
		return
	}
}
