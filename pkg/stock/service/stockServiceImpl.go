package service

import (
	"github.com/kizmey/order_management_system/entities"
	_StockRepository "github.com/kizmey/order_management_system/pkg/stock/repository"
)

type stockServiceImpl struct {
	stockRepository _StockRepository.StockRepository
}

func NewStockServiceImpl(stockRepository _StockRepository.StockRepository) StockService {
	return &stockServiceImpl{stockRepository: stockRepository}
}

func (s *stockServiceImpl) Create(stock *entities.Stock) (*entities.Stock, error) {
	return s.stockRepository.Create(stock)
}

func (s *stockServiceImpl) FindAll() (*[]entities.Stock, error) {
	return s.stockRepository.FindAll()
}

func (s *stockServiceImpl) CheckStockByProductId(id uint64) (*entities.Stock, error) {
	return s.stockRepository.CheckStockByProductId(id)
}

func (s *stockServiceImpl) Update(id uint64, stock *entities.Stock) (*entities.Stock, error) {
	return s.stockRepository.Update(id, stock)
}

func (s *stockServiceImpl) Delete(id uint64) error {
	return s.stockRepository.Delete(id)
}
