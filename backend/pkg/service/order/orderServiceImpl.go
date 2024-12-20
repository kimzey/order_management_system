package order

import (
	"context"
	"errors"
	"github.com/kizmey/order_management_system/pkg/interface/entities"
	_orderRepository "github.com/kizmey/order_management_system/pkg/repository/order"
	_stockRepository "github.com/kizmey/order_management_system/pkg/repository/stock"
	_transactionRepository "github.com/kizmey/order_management_system/pkg/repository/transaction"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type orderServiceImpl struct {
	orderRepository       _orderRepository.OrderRepository
	transactionRepository _transactionRepository.TransactionRepository
	stockRepository       _stockRepository.StockRepository
}

func NewOrderServiceImpl(orderRepository _orderRepository.OrderRepository,
	transactionRepository _transactionRepository.TransactionRepository,
	stockRepository _stockRepository.StockRepository,
) OrderService {
	return &orderServiceImpl{
		orderRepository:       orderRepository,
		transactionRepository: transactionRepository,
		stockRepository:       stockRepository,
	}
}

func (s *orderServiceImpl) Create(ctx context.Context, order *entities.Order) (*entities.Order, error) {
	ctx, sp := tracer.Start(ctx, "orderCreateService")
	defer sp.End()

	newOrder, err := s.orderRepository.Create(ctx, order)
	if err != nil {
		return nil, err
	}

	ecommerce, err := s.transactionRepository.FindProductsByTransactionID(ctx, order.TransactionID)
	if err != nil {
		return nil, err
	}

	if ecommerce.Quantity == nil {
		return nil, errors.New("quantity is nil")
	}

	stockRollback := make([]entities.Stock, 0)
	for i, product := range ecommerce.Product {
		quantityProduct := ecommerce.Quantity[i]

		stock, err := s.stockRepository.CheckStockByProductId(ctx, product.ProductID)
		if err != nil {
			return nil, err
		}

		stock.Quantity -= quantityProduct

		stock, err = s.stockRepository.Update(ctx, stock.StockID, stock)
		if err != nil {
			//for rollback
			_, _ = s.orderRepository.Delete(ctx, newOrder.OrderID)
			for _, rollback := range stockRollback {
				_, _ = s.stockRepository.Update(ctx, rollback.StockID, &rollback)
			}
			return nil, err
		}
		stock.Quantity += quantityProduct
		stockRollback = append(stockRollback, *stock)
	}

	SetOrderSubAttributes(newOrder, sp)
	return newOrder, nil
}

func (s *orderServiceImpl) FindAll(ctx context.Context) (*[]entities.Order, error) {
	ctx, sp := tracer.Start(ctx, "orderFindAllService")
	defer sp.End()

	orders, err := s.orderRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	SetOrderSubAttributes(orders, sp)
	return orders, nil
}

func (s *orderServiceImpl) FindByID(ctx context.Context, id string) (*entities.Order, error) {
	ctx, sp := tracer.Start(ctx, "orderFindByIdService")
	defer sp.End()

	order, err := s.orderRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	SetOrderSubAttributes(order, sp)
	return order, nil
}

func (s *orderServiceImpl) Update(ctx context.Context, id string, order *entities.Order) (*entities.Order, error) {
	ctx, sp := tracer.Start(ctx, "orderUpdateService")
	defer sp.End()

	newOrder, err := s.orderRepository.Update(ctx, id, order)
	if err != nil {
		return nil, err
	}

	ecommerce, err := s.transactionRepository.FindProductsByTransactionID(ctx, order.TransactionID)
	if err != nil {
		return nil, err
	}

	if ecommerce.Quantity == nil {
		return nil, errors.New("quantity is nil")
	}

	stockRollback := make([]entities.Stock, 0)
	for i, product := range ecommerce.Product {
		quantityProduct := ecommerce.Quantity[i]

		stock, err := s.stockRepository.CheckStockByProductId(ctx, product.ProductID)
		if err != nil {
			return nil, err
		}

		stock.Quantity -= quantityProduct

		stock, err = s.stockRepository.Update(ctx, stock.StockID, stock)
		if err != nil {
			_, _ = s.orderRepository.Delete(ctx, newOrder.OrderID)
			for _, rollback := range stockRollback {
				_, _ = s.stockRepository.Update(ctx, rollback.StockID, &rollback)
			}
			return nil, err
		}
		stock.Quantity += quantityProduct
		stockRollback = append(stockRollback, *stock)
	}

	SetOrderSubAttributes(newOrder, sp)
	return newOrder, nil
}

func (s *orderServiceImpl) Delete(ctx context.Context, id string) (*entities.Order, error) {
	ctx, sp := tracer.Start(ctx, "orderDeleteService")
	defer sp.End()

	order, err := s.orderRepository.Delete(ctx, id)
	if err != nil {
		return nil, err
	}
	SetOrderSubAttributes(order, sp)

	return order, err
}

func (s *orderServiceImpl) ChangeStatusNext(ctx context.Context, id string) (*entities.Order, error) {
	ctx, sp := tracer.Start(ctx, "orderChangeStatusNextService")
	defer sp.End()

	order, err := s.orderRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	err = order.NextStatus()
	if err != nil {
		return nil, err
	}

	_, err = s.orderRepository.UpdateStatus(ctx, id, order)
	if err != nil {
		return nil, err
	}
	SetOrderSubAttributes(order, sp)

	return order, nil
}
func (s *orderServiceImpl) ChageStatusDone(ctx context.Context, id string) (*entities.Order, error) {
	ctx, sp := tracer.Start(ctx, "orderChageStatusDoneService")
	defer sp.End()

	order, err := s.orderRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	err = order.NextPaidToDone()
	if err != nil {
		return nil, err
	}

	_, err = s.orderRepository.UpdateStatus(ctx, id, order)
	if err != nil {
		return nil, err
	}
	SetOrderSubAttributes(order, sp)

	return order, nil
}

func SetOrderSubAttributes(orderData any, sp trace.Span) {
	if orders, ok := orderData.(*[]entities.Order); ok {
		orderIDs := make([]string, len(*orders))
		transactionIDs := make([]string, len(*orders))
		statuses := make([]string, len(*orders))

		for _, order := range *orders {
			orderIDs = append(orderIDs, order.OrderID)
			transactionIDs = append(transactionIDs, order.TransactionID)
			statuses = append(statuses, order.Status)
		}

		sp.SetAttributes(
			attribute.StringSlice("OrderID", orderIDs),
			attribute.StringSlice("TransactionID", transactionIDs),
			attribute.StringSlice("Status", statuses),
		)
	} else if order, ok := orderData.(*entities.Order); ok {
		sp.SetAttributes(
			attribute.String("OrderID", order.OrderID),
			attribute.String("TransactionID", order.TransactionID),
			attribute.String("Status", order.Status),
		)
	} else {
		sp.RecordError(errors.New("invalid type"))
	}
}
