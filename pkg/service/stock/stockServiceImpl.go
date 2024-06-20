package stock

import (
	"context"
	"github.com/kizmey/order_management_system/pkg/interface/entities"
	"github.com/kizmey/order_management_system/pkg/interface/modelReq"
	"github.com/kizmey/order_management_system/pkg/interface/modelRes"
	_StockRepository "github.com/kizmey/order_management_system/pkg/repository/stock"
)

type stockServiceImpl struct {
	stockRepository _StockRepository.StockRepository
}

func NewStockServiceImpl(stockRepository _StockRepository.StockRepository) StockService {
	return &stockServiceImpl{stockRepository: stockRepository}
}

func (s *stockServiceImpl) Create(ctx context.Context, stock *modelReq.Stock) (*modelRes.Stock, error) {
	ctx, sp := tracer.Start(ctx, "stockCreateService")
	defer sp.End()

	stockEntity := s.stockReqToEntity(stock)
	stockEntity, err := s.stockRepository.Create(ctx, stockEntity)
	if err != nil {
		return nil, err
	}
	return s.stockEntityToRes(stockEntity), nil
}

func (s *stockServiceImpl) FindAll(ctx context.Context) (*[]modelRes.Stock, error) {
	ctx, sp := tracer.Start(ctx, "stockFindAllService")
	defer sp.End()

	stockEntities, err := s.stockRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	var stockRes []modelRes.Stock
	for _, stock := range *stockEntities {
		stockRes = append(stockRes, *s.stockEntityToRes(&stock))
	}
	return &stockRes, nil
}

func (s *stockServiceImpl) CheckStockByProductId(ctx context.Context, id string) (*modelRes.Stock, error) {
	ctx, sp := tracer.Start(ctx, "stockCheckStockByProductIdService")
	defer sp.End()

	stock, err := s.stockRepository.CheckStockByProductId(ctx, id)
	if err != nil {
		return nil, err
	}
	return s.stockEntityToRes(stock), nil
}

func (s *stockServiceImpl) Update(ctx context.Context, id string, stock *modelReq.Stock) (*modelRes.Stock, error) {
	ctx, sp := tracer.Start(ctx, "stockUpdateService")
	defer sp.End()

	stockEntity := s.stockReqToEntity(stock)
	stockEntity, err := s.stockRepository.Update(ctx, id, stockEntity)
	if err != nil {
		return nil, err
	}
	return s.stockEntityToRes(stockEntity), nil
}

func (s *stockServiceImpl) Delete(ctx context.Context, id string) (*modelRes.Stock, error) {
	ctx, sp := tracer.Start(ctx, "stockDeleteService")
	defer sp.End()

	stock, err := s.stockRepository.Delete(ctx, id)
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
