package service

import (
	"errors"
	"github.com/kizmey/order_management_system/pkg/interface/entities"
	"github.com/kizmey/order_management_system/pkg/interface/modelReq"
	"github.com/kizmey/order_management_system/pkg/interface/modelRes"
	_transactionRepository "github.com/kizmey/order_management_system/pkg/repository"
)

type orderServiceImpl struct {
	orderRepository       _transactionRepository.OrderRepository
	transactionRepository _transactionRepository.TransactionRepository
	stockRepository       _transactionRepository.StockRepository
	productRepository     _transactionRepository.ProductRepository
}

func NewOrderServiceImpl(orderRepository _transactionRepository.OrderRepository,
	transactionRepository _transactionRepository.TransactionRepository,
	stockRepository _transactionRepository.StockRepository,
) OrderService {
	return &orderServiceImpl{
		orderRepository:       orderRepository,
		transactionRepository: transactionRepository,
		stockRepository:       stockRepository,
	}
}

func (s *orderServiceImpl) Create(order *modelReq.Order) (*modelRes.Order, error) {
	ecommerce, err := s.transactionRepository.FindProductsByTransactionID(order.TransactionID)
	if err != nil {
		return nil, err
	}

	if ecommerce.Quantity == nil {
		return nil, errors.New("quantity is nil")
	}

	orderEntity := s.orderReqToEntity(order)
	newOrder, err := s.orderRepository.Create(orderEntity)
	if err != nil {
		return nil, err
	}

	for i, product := range ecommerce.Product {
		quantityProduct := ecommerce.Quantity[i]

		stock, err := s.stockRepository.CheckStockByProductId(product.ProductID)
		if err != nil {
			return nil, err
		}

		if stock.Quantity < quantityProduct {
			return nil, errors.New("stock not enough")
		}
		stock.Quantity -= quantityProduct

		stock, err = s.stockRepository.Update(stock.StockID, stock)
		if err != nil {
			_, _ = s.orderRepository.Delete(newOrder.OrderID)
			return nil, err
		}
	}

	return s.orderEntityToModelRes(newOrder), nil
}

func (s *orderServiceImpl) FindAll() (*[]modelRes.Order, error) {
	orders, err := s.orderRepository.FindAll()
	if err != nil {
		return nil, err
	}

	allOrder := make([]modelRes.Order, 0)
	for _, order := range *orders {
		allOrder = append(allOrder, *s.orderEntityToModelRes(&order))
	}
	return &allOrder, nil
}

func (s *orderServiceImpl) FindByID(id string) (*modelRes.Order, error) {

	order, err := s.orderRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return s.orderEntityToModelRes(order), nil
}

func (s *orderServiceImpl) Update(id string, order *modelReq.Order) (*modelRes.Order, error) {
	ecommerce, err := s.transactionRepository.FindProductsByTransactionID(order.TransactionID)
	if err != nil {
		return nil, err
	}

	orderEntity := s.orderReqToEntity(order)

	if ecommerce.Quantity == nil {
		return nil, errors.New("quantity is nil")
	}

	newOrder, err := s.orderRepository.Update(id, orderEntity)
	if err != nil {
		return nil, err
	}

	for i, product := range ecommerce.Product {
		quantityProduct := ecommerce.Quantity[i]

		stock, err := s.stockRepository.CheckStockByProductId(product.ProductID)
		if err != nil {
			return nil, err
		}

		if stock.Quantity < quantityProduct {
			return nil, errors.New("stock not enough")
		}
		stock.Quantity -= quantityProduct

		stock, err = s.stockRepository.Update(stock.StockID, stock)
		if err != nil {
			_, _ = s.orderRepository.Delete(newOrder.OrderID)
			return nil, err
		}
	}

	return s.orderEntityToModelRes(newOrder), nil
}

func (s *orderServiceImpl) Delete(id string) (*modelRes.Order, error) {

	order, err := s.orderRepository.Delete(id)
	if err != nil {
		return nil, err
	}
	return s.orderEntityToModelRes(order), err
}

func (s *orderServiceImpl) ChangeStatusNext(id string) (*modelRes.Order, error) {

	order, err := s.orderRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	err = order.NextStatus()
	if err != nil {
		return nil, err
	}

	_, err = s.orderRepository.UpdateStatus(id, order)
	if err != nil {
		return nil, err
	}

	return s.orderEntityToModelRes(order), nil
}
func (s *orderServiceImpl) ChageStatusDone(id string) (*modelRes.Order, error) {

	order, err := s.orderRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	err = order.NextPaidToDone()
	if err != nil {
		return nil, err
	}

	_, err = s.orderRepository.UpdateStatus(id, order)
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
