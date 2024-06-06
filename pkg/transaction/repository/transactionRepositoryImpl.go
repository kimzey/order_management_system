package repository

import (
	"github.com/kizmey/order_management_system/database"
	"github.com/kizmey/order_management_system/entities"
	"github.com/kizmey/order_management_system/model"
	"github.com/labstack/echo/v4"
)

type transactionRepositoryImpl struct {
	db     database.Database
	logger echo.Logger
}

func NewTransactionController(db database.Database, logger echo.Logger) TransactionRepository {
	return &transactionRepositoryImpl{db: db, logger: logger}
}

func (r *transactionRepositoryImpl) Create(transaction *entities.Transaction) (uint64, error) {
	newTransaction := new(model.Transaction)

	if err := r.db.Connect().Create(transaction.ToTransactionModel()).Scan(&newTransaction).Error; err != nil {
		r.logger.Error("Creating item failed:", err.Error())
		return 0, err
	}
	return newTransaction.ID, nil
}
