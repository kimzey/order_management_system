package repository

import (
	"github.com/kizmey/order_management_system/database"
	"github.com/kizmey/order_management_system/entities"
	"github.com/kizmey/order_management_system/model"
)

type transactionRepositoryImpl struct {
	db database.Database
}

func NewTransactionController(db database.Database) TransactionRepository {
	return &transactionRepositoryImpl{db: db}
}

func (r *transactionRepositoryImpl) Create(transaction *entities.Transaction) (*entities.Transaction, error) {
	//check stock ก่อน

	//stock := new(model.Stock)
	//if err := r.db.Connect().Where("product_id = ? AND quantity >= ?", transaction.ProductID, transaction.Quantity).First(&stock).Error; err != nil {
	//	return 0, err
	//}

	transactionModel := ToTransactionModel(transaction)
	if err := r.db.Connect().Create(transactionModel).Preload("Product").Error; err != nil {
		return nil, err
	}

	if err := r.db.Connect().Preload("Product").First(&transactionModel, transactionModel.ID).Error; err != nil {
		return nil, err
	}
	return transactionModel.ToTransactionEntity(), nil
}

func (r *transactionRepositoryImpl) FindAll() (*[]entities.Transaction, error) {
	transactions := new([]model.Transaction)

	if err := r.db.Connect().Preload("Product").Find(&transactions).Error; err != nil {
		return nil, err
	}
	allTransactions := model.ConvertModelsTransactionToEntities(transactions)
	return allTransactions, nil
}

func (r *transactionRepositoryImpl) FindByID(id uint64) (*entities.Transaction, error) {

	transaction := new(model.Transaction)
	if err := r.db.Connect().Preload("Product").Where("id = ?", id).First(&transaction, id).Error; err != nil {
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
		return nil, err
	}

	if err := r.db.Connect().Preload("Product").First(&transactionModel, id).Error; err != nil {
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
