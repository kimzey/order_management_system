package service

import (
	"github.com/kizmey/order_management_system/entities"
	_TransactionRepository "github.com/kizmey/order_management_system/pkg/transaction/repository"
)

type transactionService struct {
	transactionRepository _TransactionRepository.TransactionRepository
}

func NewTransactionService(transactionRepository _TransactionRepository.TransactionRepository) TransactionService {
	return &transactionService{transactionRepository}
}

func (s *transactionService) Create(transaction *entities.Transaction) (*entities.Transaction, error) {
	_, err := s.transactionRepository.Create(transaction)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *transactionService) FindAll() (*[]entities.Transaction, error) {
	return s.transactionRepository.FindAll()
}
