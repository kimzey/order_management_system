package repository

import (
	"errors"
	"fmt"
	"github.com/kizmey/order_management_system/database"
	"github.com/kizmey/order_management_system/pkg/interface/entities"
	"github.com/kizmey/order_management_system/pkg/interface/model"
)

type transactionRepositoryImpl struct {
	db database.Database
}

func NewTransactionController(db database.Database) TransactionRepository {
	return &transactionRepositoryImpl{db: db}
}

func (r *transactionRepositoryImpl) Create(transaction *entities.Transaction) (*entities.Transaction, error) {
	//var product = new(model.Product)
	//var stock = new(model.Stock)
	//
	//err := r.db.Connect().Joins("JOIN stocks ON stocks.product_id = products.id").
	//	Where("products.id = ? AND stocks.quantity >= ?", transaction.ProductID, transaction.Quantity).
	//	First(&product).First(&stock).Error
	//if err != nil {
	//	return nil, errors.New("id not correct or not enough stock")
	//}
	//
	//transaction.SumPrice = transaction.CalculatePrice(product.Price, transaction.Quantity, transaction.IsDomestic)
	transactionModel := ToTransactionModel(transaction)

	if err := r.db.Connect().Create(transactionModel).Preload("Product").Preload("Product").First(&transactionModel, transactionModel.ID).Error; err != nil {
		return nil, errors.New(fmt.Sprintf("failed to create transaction: %s", err))
	}

	return transactionModel.ToTransactionEntity(), nil
}

func (r *transactionRepositoryImpl) FindAll() (*[]entities.Transaction, error) {
	transactions := new([]model.Transaction)

	if err := r.db.Connect().Preload("Product").Find(&transactions).Error; err != nil {
		return nil, errors.New(fmt.Sprintf("failed to find transactions: %s", err))
	}
	allTransactions := model.ConvertModelsTransactionToEntities(transactions)
	return allTransactions, nil
}

func (r *transactionRepositoryImpl) FindByID(id uint64) (*entities.Transaction, error) {

	transaction := new(model.Transaction)
	if err := r.db.Connect().Preload("Product").Where("id = ?", id).First(&transaction, id).Error; err != nil {
		return nil, errors.New(fmt.Sprintf("failed to find transaction: %s", err))
	}

	return transaction.ToTransactionEntity(), nil
}
func (r *transactionRepositoryImpl) Update(id uint64, transaction *entities.Transaction) (*entities.Transaction, error) {
	transactionModel := ToTransactionModel(transaction)

	if err := r.db.Connect().Model(&transactionModel).Where(
		"id = ?", id,
	).Updates(
		transactionModel,
	).Scan(transactionModel).Preload("Product").First(&transactionModel, id).Error; err != nil {
		return nil, errors.New(fmt.Sprintf("failed to update transaction: %s", err))
	}
	return transactionModel.ToTransactionEntity(), nil
}

func (r *transactionRepositoryImpl) Delete(id uint64) error {

	err := r.db.Connect().Delete(&model.Transaction{}, id).Error
	if err != nil {
		return errors.New(fmt.Sprintf("failed to delete transaction: %s", err))
	}
	fmt.Println(err)
	return nil

}

func ToTransactionModel(e *entities.Transaction) *model.Transaction {
	return &model.Transaction{
		ProductID:  e.ProductID,
		Quantity:   e.Quantity,
		SumPrice:   e.SumPrice,
		IsDomestic: e.IsDomestic,
	}
}
