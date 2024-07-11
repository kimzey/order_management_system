package stock

import (
	"context"
	"github.com/kizmey/order_management_system/pkg/interface/modelRes"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"

	"errors"
	"fmt"
	"github.com/kizmey/order_management_system/database"
	"github.com/kizmey/order_management_system/pkg/interface/entities"
	"github.com/kizmey/order_management_system/pkg/interface/model"
)

type stockRepositoryImpl struct {
	db database.Database
}

func NewStockRepositoryImpl(db database.Database) StockRepository {
	return &stockRepositoryImpl{db: db}
}

func (r *stockRepositoryImpl) Create(ctx context.Context, stock *entities.Stock) (*entities.Stock, error) {
	_, sp := tracer.Start(ctx, "stockCreateRepository")
	defer sp.End()

	modelStock := r.ToStockModel(stock)
	newStock := new(model.Stock)

	if err := r.db.Connect().Create(modelStock).Scan(newStock).Error; err != nil {
		return nil, errors.New(fmt.Sprintf("failed to create stock"))
	}

	stockEntity := newStock.ToStockEntity()
	r.SetStockSubAttributes(stockEntity, sp)
	return stockEntity, nil
}

func (r *stockRepositoryImpl) FindAll(ctx context.Context) (*[]entities.Stock, error) {
	_, sp := tracer.Start(ctx, "stockFindAllRepository")
	defer sp.End()

	stocks := new([]model.Stock)
	if err := r.db.Connect().Find(stocks).Error; err != nil {
		return nil, errors.New(fmt.Sprintf("failed to find all stock"))
	}

	allStock := model.ConvertStockModelsToEntities(stocks)

	r.SetStockSubAttributes(allStock, sp)
	return allStock, nil
}

func (r *stockRepositoryImpl) CheckStockByProductId(ctx context.Context, productId string) (*entities.Stock, error) {
	_, sp := tracer.Start(ctx, "stockCheckStockByProductIdRepository")
	defer sp.End()

	stock := new(model.Stock)
	if err := r.db.Connect().Preload("Product").Where("product_id = ?", productId).First(stock).Error; err != nil {
		return nil, errors.New(fmt.Sprintf("failed to find stock"))
	}

	stockEntity := stock.ToStockEntity()
	r.SetStockSubAttributes(stockEntity, sp)
	return stockEntity, nil
}

func (r *stockRepositoryImpl) Update(ctx context.Context, stockid string, stock *entities.Stock) (*entities.Stock, error) {
	_, sp := tracer.Start(ctx, "stockUpdateRepository")
	defer sp.End()

	stocks := new(model.Stock)
	modelStock := r.ToStockModel(stock)

	if modelStock.Quantity == 0 {
		if err := r.db.Connect().Model(&modelStock).
			Where("id = ?", stockid).
			Update("quantity", modelStock.Quantity).
			Scan(stocks).Error; err != nil {
			return nil, errors.New(fmt.Sprintf("failed to update stock"))
		}
	} else {
		if err := r.db.Connect().Model(&modelStock).
			Where("id = ? AND ? >= 0", stockid, modelStock.Quantity).
			Update("quantity", modelStock.Quantity).
			Scan(stocks).Error; err != nil {
			return nil, errors.New(fmt.Sprintf("failed to update stock"))
		}
	}

	stockEntity := stocks.ToStockEntity()
	r.SetStockSubAttributes(stockEntity, sp)
	return stockEntity, nil
}

func (r *stockRepositoryImpl) Delete(ctx context.Context, id string) (*entities.Stock, error) {
	_, sp := tracer.Start(ctx, "stockDeleteRepository")
	defer sp.End()

	stock := new(model.Stock)
	if err := r.db.Connect().Where("id = ?", id).First(&stock).Delete(&stock).Error; err != nil {
		return nil, errors.New(fmt.Sprintf("failed to delete stock"))
	}

	stockEntity := stock.ToStockEntity()
	r.SetStockSubAttributes(stockEntity, sp)
	return stockEntity, nil
}

func (r *stockRepositoryImpl) ToStockModel(e *entities.Stock) *model.Stock {
	return &model.Stock{
		ProductID: e.ProductID,
		Quantity:  e.Quantity,
	}
}

func (r *stockRepositoryImpl) ToStockModelRes(e *entities.Stock) *modelRes.Stock {
	return &modelRes.Stock{
		StockID:   e.StockID,
		ProductID: e.ProductID,
		Quantity:  e.Quantity,
	}
}

func (r *stockRepositoryImpl) SetStockSubAttributes(stockData any, sp trace.Span) {
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
