package repository

import (
	"errors"
	"fmt"
	"github.com/kizmey/order_management_system/database"
	_interface "github.com/kizmey/order_management_system/pkg/interface"
	"github.com/kizmey/order_management_system/pkg/interface/entities"
	"github.com/kizmey/order_management_system/pkg/interface/model"
)

type transactionRepositoryImpl struct {
	db database.Database
}

func NewTransactionRepositoryImpl(db database.Database) TransactionRepository {
	return &transactionRepositoryImpl{db: db}
}

func (r *transactionRepositoryImpl) Create(transaction *_interface.TransactionEcommerce) (*entities.Transaction, error) {
	transactionModel := ToTransactionModel(transaction)

	if err := r.db.Connect().Create(&transactionModel).Error; err != nil {
		return nil, errors.New(fmt.Sprintf("failed to create transaction: %s", err))
	}

	fmt.Println("_______________________")
	fmt.Println(transactionModel)

	for productID, quantity := range transaction.AddessProduct {
		if err := r.db.Connect().Model(&model.TransactionProduct{}).Where("transaction_id = ? AND product_id = ?", transactionModel.ID, productID).Update("quantity", quantity).Error; err != nil {
			return nil, errors.New(fmt.Sprintf("failed to update transaction: %s", err))
		}
	}

	return transactionModel.ToTransactionEntity(), nil
}

func (r *transactionRepositoryImpl) FindAll() (*[]entities.Transaction, error) {
	transactions := new([]model.Transaction)

	if err := r.db.Connect().Find(&transactions).Error; err != nil {
		return nil, errors.New(fmt.Sprintf("failed to find transactions: %s", err))
	}
	allTransactions := model.ConvertModelsTransactionToEntities(transactions)
	return allTransactions, nil
}

func (r *transactionRepositoryImpl) FindByID(id string) (*entities.Transaction, error) {

	transaction := new(model.Transaction)
	if err := r.db.Connect().Preload("Product").Where("id = ?", id).First(&transaction).Error; err != nil {
		return nil, errors.New(fmt.Sprintf("failed to find transaction: %s", err))
	}

	return transaction.ToTransactionEntity(), nil
}
func (r *transactionRepositoryImpl) Update(id string, transaction *_interface.TransactionEcommerce) (*entities.Transaction, error) {
	transactionModel := ToTransactionModel(transaction)

	if err := r.db.Connect().Model(&transactionModel).Where("id = ?", id).Updates(&transactionModel); err != nil {
		return nil, errors.New(fmt.Sprintf("failed to create transaction: %s", err))
	}

	for productID, quantity := range transaction.AddessProduct {
		if err := r.db.Connect().Model(&model.TransactionProduct{}).Where("transaction_id = ? AND product_id = ?", transactionModel.ID, productID).Update("quantity", quantity).Error; err != nil {
			return nil, errors.New(fmt.Sprintf("failed to update transaction: %s", err))
		}
	}

	return transactionModel.ToTransactionEntity(), nil

	//transactionModel := ToTransactionModel(transaction)
	//
	//if err := r.db.Connect().Model(&transactionModel).Where(
	//	"id = ?", id,
	//).Updates(
	//	transactionModel,
	//).Scan(transactionModel).Preload("Product").First(&transactionModel).Where("id = ?", id).Error; err != nil {
	//	return nil, errors.New(fmt.Sprintf("failed to update transaction: %s", err))
	//}
	//return transactionModel.ToTransactionEntity(), nil

}

func (r *transactionRepositoryImpl) Delete(id string) error {

	err := r.db.Connect().Where("id = ?", id).Delete(&model.Transaction{}).Error
	if err != nil {
		return errors.New(fmt.Sprintf("failed to delete transaction: %s", err))
	}
	fmt.Println(err)
	return nil

}

func ToTransactionModel(e *_interface.TransactionEcommerce) *model.Transaction {
	var productlist []model.Product
	for _, v := range e.Product {
		productlist = append(productlist, model.Product{
			ID:    v.ProductID,
			Name:  v.ProductName,
			Price: v.ProductPrice,
		})
	}
	return &model.Transaction{
		SumPrice:   e.Tranasaction.SumPrice,
		IsDomestic: e.Tranasaction.IsDomestic,
		Products:   productlist,
	}
}
