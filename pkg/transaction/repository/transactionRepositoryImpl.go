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

	if err := r.db.Connect().Preload("Product").Find(&transactions).Error; err != nil {
		r.logger.Error("Failed to find transactions:", err.Error())
		return nil, err
	}
	allTransactions := model.ConvertModelsTransactionToEntities(transactions)
	return allTransactions, nil
}

func (r *transactionRepositoryImpl) FindByID(id uint64) (*entities.Transaction, error) {

	transaction := new(model.Transaction)
	if err := r.db.Connect().Preload("Product").Where("id = ?", id).First(&transaction, id).Error; err != nil {
		r.logger.Error("Failed to find transaction:", err.Error())
		return nil, err
	}

	return transaction.ToTransactionEntity(), nil
}
func (r *transactionRepositoryImpl) Update(id uint64, transaction *entities.Transaction) (*entities.Transaction, error) {
	transactionModel := ToTransactionModel(transaction)

	if err := r.db.Connect().Model(&transactionModel).Where(
		"id = ?", id,
	).Updates(
		transactionModel,
	).Scan(transactionModel).Error; err != nil {
		r.logger.Error("Editing item failed:", err.Error())
		return nil, err
	}

	if err := r.db.Connect().Preload("Product").First(&transactionModel, id).Error; err != nil {
		r.logger.Error("Failed to find transaction:", err.Error())
		return nil, err
	}

	return transactionModel.ToTransactionEntity(), nil
}

func (r *transactionRepositoryImpl) Delete(id uint64) error {
	return r.db.Connect().Delete(&model.Transaction{}, id).Error

}

func ToTransactionModel(e *entities.Transaction) *model.Transaction {
	return &model.Transaction{
		ProductID:  e.ProductID,
		Quantity:   e.Quantity,
		SumPrice:   e.SumPrice,
		IsDomestic: e.IsDomestic,
	}
}
