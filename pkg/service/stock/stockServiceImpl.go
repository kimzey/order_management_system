package stock

import (
	"context"
	"errors"
	"github.com/kizmey/order_management_system/pkg/interface/entities"
	_StockRepository "github.com/kizmey/order_management_system/pkg/repository/stock"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
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

func (s *stockServiceImpl) SetStockSubAttributes(stockData any, sp trace.Span) {
	if stocks, ok := stockData.(*[]entities.Stock); ok {
		stockIDs := make([]string, len(*stocks))
		productIDs := make([]string, len(*stocks))
		quantities := make([]int, len(*stocks))

		for _, stock := range *stocks {
			stockIDs = append(stockIDs, stock.StockID)
			productIDs = append(productIDs, stock.ProductID)
			quantities = append(quantities, int(stock.Quantity))
		}

		sp.SetAttributes(
			attribute.StringSlice("StockID", stockIDs),
			attribute.StringSlice("ProductID", productIDs),
			attribute.IntSlice("Quantity", quantities),
		)
	} else if stock, ok := stockData.(*entities.Stock); ok {
		sp.SetAttributes(
			attribute.String("StockID", stock.StockID),
			attribute.String("ProductID", stock.ProductID),
			attribute.Int("Quantity", int(stock.Quantity)),
		)
	} else {
		sp.RecordError(errors.New("invalid type"))
	}
}
