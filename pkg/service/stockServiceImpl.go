package service

import (
	"github.com/kizmey/order_management_system/pkg/interface/entities"
	"github.com/kizmey/order_management_system/pkg/interface/modelReq"
	"github.com/kizmey/order_management_system/pkg/interface/modelRes"
	_StockRepository "github.com/kizmey/order_management_system/pkg/repository"
)

type stockServiceImpl struct {
	stockRepository _StockRepository.StockRepository
}

func NewStockServiceImpl(stockRepository _StockRepository.StockRepository) StockService {
	return &stockServiceImpl{stockRepository: stockRepository}
}

func (s *stockServiceImpl) Create(stock *modelReq.Stock) (*modelRes.Stock, error) {

	stockEntity := s.stockReqToEntity(stock)
	stockEntity, err := s.stockRepository.Create(stockEntity)
	if err != nil {
		return nil, err
	}
	return s.stockEntityToRes(stockEntity), nil
}

func (s *stockServiceImpl) FindAll() (*[]modelRes.Stock, error) {

	stockEntities, err := s.stockRepository.FindAll()
	if err != nil {
		return nil, err
	}
	var stockRes []modelRes.Stock
	for _, stock := range *stockEntities {
		stockRes = append(stockRes, *s.stockEntityToRes(&stock))
	}
	return &stockRes, nil
}

func (s *stockServiceImpl) CheckStockByProductId(id string) (*modelRes.Stock, error) {

	stock, err := s.stockRepository.CheckStockByProductId(id)
	if err != nil {
		return nil, err
	}
	return s.stockEntityToRes(stock), nil
}

func (s *stockServiceImpl) Update(id string, stock *modelReq.Stock) (*modelRes.Stock, error) {

	stockEntity := s.stockReqToEntity(stock)
	stockEntity, err := s.stockRepository.Update(id, stockEntity)
	if err != nil {
		return nil, err
	}
	return s.stockEntityToRes(stockEntity), nil
}

func (s *stockServiceImpl) Delete(id string) (*modelRes.Stock, error) {

	stock, err := s.stockRepository.Delete(id)
	if err != nil {
		return nil, err
	}

	return s.stockEntityToRes(stock), nil
}

func (s *stockServiceImpl) stockReqToEntity(stockReq *modelReq.Stock) *entities.Stock {
	return &entities.Stock{
		ProductID: stockReq.ProductID,
		Quantity:  stockReq.Quantity,
	}
}

func (s *stockServiceImpl) stockEntityToRes(stock *entities.Stock) *modelRes.Stock {
	return &modelRes.Stock{
		StockID:   stock.StockID,
		ProductID: stock.ProductID,
		Quantity:  stock.Quantity,
	}
}
