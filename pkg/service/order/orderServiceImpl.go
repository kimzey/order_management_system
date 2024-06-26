package order

import (
	"context"
	"errors"
	"github.com/kizmey/order_management_system/pkg/interface/entities"
	"github.com/kizmey/order_management_system/pkg/interface/modelReq"
	"github.com/kizmey/order_management_system/pkg/interface/modelRes"
	_orderRepository "github.com/kizmey/order_management_system/pkg/repository/order"
	_stockRepository "github.com/kizmey/order_management_system/pkg/repository/stock"
	_transactionRepository "github.com/kizmey/order_management_system/pkg/repository/transaction"
	customTracer "github.com/kizmey/order_management_system/tracer"
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

func (s *orderServiceImpl) Create(ctx context.Context, order *modelReq.Order) (*modelRes.Order, error) {
	ctx, sp := tracer.Start(ctx, "orderCreateService")
	defer sp.End()

	ecommerce, err := s.transactionRepository.FindProductsByTransactionID(ctx, order.TransactionID)
	if err != nil {
		return nil, err
	}

	if ecommerce.Quantity == nil {
		return nil, errors.New("quantity is nil")
	}

	orderEntity := s.orderReqToEntity(order)
	newOrder, err := s.orderRepository.Create(ctx, orderEntity)
	if err != nil {
		return nil, err
	}

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
			return nil, err
		}
	}

	customTracer.SetSubAttributesWithJson(newOrder, sp)

	return s.orderEntityToModelRes(newOrder), nil
}

func (s *orderServiceImpl) FindAll(ctx context.Context) (*[]modelRes.Order, error) {
	ctx, sp := tracer.Start(ctx, "orderFindAllService")
	defer sp.End()

	orders, err := s.orderRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	allOrder := make([]modelRes.Order, 0)
	for _, order := range *orders {
		allOrder = append(allOrder, *s.orderEntityToModelRes(&order))
	}
	return &allOrder, nil
}

func (s *orderServiceImpl) FindByID(ctx context.Context, id string) (*modelRes.Order, error) {
	ctx, sp := tracer.Start(ctx, "orderFindByIdService")
	defer sp.End()

	order, err := s.orderRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return s.orderEntityToModelRes(order), nil
}

func (s *orderServiceImpl) Update(ctx context.Context, id string, order *modelReq.Order) (*modelRes.Order, error) {
	ctx, sp := tracer.Start(ctx, "orderUpdateService")
	defer sp.End()

	ecommerce, err := s.transactionRepository.FindProductsByTransactionID(ctx, order.TransactionID)
	if err != nil {
		return nil, err
	}

	orderEntity := s.orderReqToEntity(order)

	if ecommerce.Quantity == nil {
		return nil, errors.New("quantity is nil")
	}

	newOrder, err := s.orderRepository.Update(ctx, id, orderEntity)
	if err != nil {
		return nil, err
	}

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
			return nil, err
		}
	}

	return s.orderEntityToModelRes(newOrder), nil
}

func (s *orderServiceImpl) Delete(ctx context.Context, id string) (*modelRes.Order, error) {
	ctx, sp := tracer.Start(ctx, "orderDeleteService")
	defer sp.End()

	order, err := s.orderRepository.Delete(ctx, id)
	if err != nil {
		return nil, err
	}
	return s.orderEntityToModelRes(order), err
}

func (s *orderServiceImpl) ChangeStatusNext(ctx context.Context, id string) (*modelRes.Order, error) {
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

	return s.orderEntityToModelRes(order), nil
}
func (s *orderServiceImpl) ChageStatusDone(ctx context.Context, id string) (*modelRes.Order, error) {
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

	return s.orderEntityToModelRes(order), nil
}

func (s *orderServiceImpl) orderReqToEntity(orderReq *modelReq.Order) *entities.Order {
	return &entities.Order{
		TransactionID: orderReq.TransactionID,
		Status:        orderReq.Status,
	}
}

func (s *orderServiceImpl) orderEntityToModelRes(order *entities.Order) *modelRes.Order {
	return &modelRes.Order{
		OrderID:       order.OrderID,
		TransactionID: order.TransactionID,
		//ProductID:     order.ProductID,
		Status: order.Status,
	}
}
