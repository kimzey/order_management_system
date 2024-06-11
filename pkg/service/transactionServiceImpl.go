package service

import (
	"github.com/kizmey/order_management_system/pkg/interface/entities"
	"github.com/kizmey/order_management_system/pkg/interface/modelReq"
	"github.com/kizmey/order_management_system/pkg/interface/modelRes"
	_TransactionRepository "github.com/kizmey/order_management_system/pkg/repository"
)

type transactionService struct {
	transactionRepository _TransactionRepository.TransactionRepository
	stockRepository       _TransactionRepository.StockRepository
	productRepository     _TransactionRepository.ProductRepository
}

func NewTransactionServiceImpl(
	transactionRepository _TransactionRepository.TransactionRepository,
	productRepository _TransactionRepository.ProductRepository,
) TransactionService {
	return &transactionService{
		transactionRepository: transactionRepository,
		productRepository:     productRepository,
	}
}

func (s *transactionService) Create(transaction *modelReq.Transaction) (*modelRes.Transaction, error) {
	transactionEntity := s.transactionReqToEntity(transaction)

	product, err := s.productRepository.FindByID(transaction.ProductID)
	if err != nil {
		return nil, err
	}
	transactionEntity.SumPrice = transactionEntity.CalculatePrice(product.ProductPrice, transaction.Quantity, transaction.IsDomestic)
	transactionEntity, err = s.transactionRepository.Create(transactionEntity)
	if err != nil {
		return nil, err
	}

	return s.transactionEntityToRes(transactionEntity), nil
}

func (s *transactionService) FindAll() (*[]modelRes.Transaction, error) {

	transactionEntities, err := s.transactionRepository.FindAll()
	if err != nil {
		return nil, err
	}

	allTransaction := make([]modelRes.Transaction, 0)
	for _, transactionEntity := range *transactionEntities {
		allTransaction = append(allTransaction, *s.transactionEntityToRes(&transactionEntity))
	}

	return &allTransaction, nil
}

func (s *transactionService) FindByID(id uint64) (*modelRes.Transaction, error) {
	transactionEntity, err := s.transactionRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return s.transactionEntityToRes(transactionEntity), nil
}

func (s *transactionService) Update(id uint64, transaction *modelReq.Transaction) (*modelRes.Transaction, error) {
	product, err := s.productRepository.FindByID(transaction.ProductID)
	if err != nil {
		return nil, err
	}

	transactionEntity := s.transactionReqToEntity(transaction)
	transactionEntity.SumPrice = transactionEntity.CalculatePrice(product.ProductPrice, transaction.Quantity, transaction.IsDomestic)

	transactionEntity, err = s.transactionRepository.Update(id, transactionEntity)
	if err != nil {
		return nil, err
	}
	return s.transactionEntityToRes(transactionEntity), nil
}

func (s *transactionService) Delete(id uint64) error {

	err := s.transactionRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (s *transactionService) transactionReqToEntity(transactionReq *modelReq.Transaction) *entities.Transaction {
	return &entities.Transaction{
		ProductID:  transactionReq.ProductID,
		Quantity:   transactionReq.Quantity,
		IsDomestic: transactionReq.IsDomestic,
	}
}

func (s *transactionService) transactionEntityToRes(transactionEntity *entities.Transaction) *modelRes.Transaction {
	return &modelRes.Transaction{
		TransactionID: transactionEntity.TransactionID,
		ProductID:     transactionEntity.ProductID,
		ProductName:   transactionEntity.ProductName,
		Quantity:      transactionEntity.Quantity,
		IsDomestic:    transactionEntity.IsDomestic,
		SumPrice:      transactionEntity.SumPrice,
	}
}
