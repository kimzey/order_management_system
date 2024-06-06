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
	modelTransaction := ToTransactionModel(transaction)
	newTransaction := new(model.Transaction)

	if err := r.db.Connect().Create(modelTransaction).Scan(&newTransaction).Error; err != nil {
		r.logger.Error("Creating item failed:", err.Error())
		return 0, err
	}
	return newTransaction.ID, nil
}

func (r *transactionRepositoryImpl) FindAll() (*[]entities.Transaction, error) {
	transactions := new([]model.Transaction)

	if err := r.db.Connect().Find(&transactions).Error; err != nil {
		r.logger.Error("Failed to find transactions:", err.Error())
		return nil, err
	}
	allTransactions := model.ConvertModelsTransactionToEntities(transactions)
	return allTransactions, nil
}

func ToTransactionModel(e *entities.Transaction) *model.Transaction {
	return &model.Transaction{
		ProductID:  e.ProductID,
		Quantity:   e.Quantity,
		SumPrice:   e.SumPrice,
		IsDomestic: e.IsDomestic,
	}
}
