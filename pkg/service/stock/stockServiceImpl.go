package stock

import (
	"context"
	"github.com/kizmey/order_management_system/pkg/interface/entities"
	_StockRepository "github.com/kizmey/order_management_system/pkg/repository/stock"
)

type stockServiceImpl struct {
	stockRepository _StockRepository.StockRepository
}

func NewStockServiceImpl(stockRepository _StockRepository.StockRepository) StockService {
	return &stockServiceImpl{stockRepository: stockRepository}
}

func (s *stockServiceImpl) Create(ctx context.Context, stock *entities.Stock) (*entities.Stock, error) {
	ctx, sp := tracer.Start(ctx, "stockCreateService")
	defer sp.End()

	stockEntity, err := s.stockRepository.Create(ctx, stock)
	if err != nil {
		return nil, err
	}
	return stockEntity, nil
}

func (s *stockServiceImpl) FindAll(ctx context.Context) (*[]entities.Stock, error) {
	ctx, sp := tracer.Start(ctx, "stockFindAllService")
	defer sp.End()

	stockEntities, err := s.stockRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return stockEntities, nil
}

func (s *stockServiceImpl) CheckStockByProductId(ctx context.Context, id string) (*entities.Stock, error) {
	ctx, sp := tracer.Start(ctx, "stockCheckStockByProductIdService")
	defer sp.End()

	stock, err := s.stockRepository.CheckStockByProductId(ctx, id)
	if err != nil {
		return nil, err
	}
	return stock, nil
}

func (s *stockServiceImpl) Update(ctx context.Context, id string, stock *entities.Stock) (*entities.Stock, error) {
	ctx, sp := tracer.Start(ctx, "stockUpdateService")
	defer sp.End()

	stockEntity, err := s.stockRepository.Update(ctx, id, stock)
	if err != nil {
		return nil, err
	}
	return stockEntity, nil
}

func (s *stockServiceImpl) Delete(ctx context.Context, id string) (*entities.Stock, error) {
	ctx, sp := tracer.Start(ctx, "stockDeleteService")
	defer sp.End()

	stock, err := s.stockRepository.Delete(ctx, id)
	if err != nil {
		return nil, err
	}

	return stock, nil
}
