package service

import (
	_TransactionRepository "github.com/kizmey/order_management_system/pkg/transaction/repository"
)

type transactionService struct {
	transactionRepository _TransactionRepository.TransactionRepository
}

func NewTransactionService(transactionRepository _TransactionRepository.TransactionRepository) TransactionService {
	return &transactionService{transactionRepository}
}
