package service

import (
	"errors"
	"github.com/kizmey/order_management_system/entities"
	"github.com/kizmey/order_management_system/modelReq"
	"github.com/kizmey/order_management_system/modelRes"
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
	productRepository _transactionRepository.ProductRepository,
) OrderService {
	return &orderServiceImpl{
		orderRepository:       orderRepository,
		transactionRepository: transactionRepository,
		stockRepository:       stockRepository,
		productRepository:     productRepository,
	}
}

func (s *orderServiceImpl) Create(order *modelReq.Order) (*modelRes.Order, error) {

	stock, err := s.stockRepository.CheckStockByProductId(order.ProductID)
	if err != nil {
		return nil, err
	}

	transaction, err := s.transactionRepository.FindByID(order.TransactionID)
	if err != nil {
		return nil, err
	}

	if stock.Quantity < transaction.Quantity {
		return nil, errors.New("stock not enough")
	}

	stock.Quantity = stock.Quantity - transaction.Quantity
	_, err = s.stockRepository.Update(stock.StockID, stock)
	if err != nil {
		return nil, err
	}

	createOrder := s.orderReqToEntity(order)
	newOrder, err := s.orderRepository.Create(createOrder)
	if err != nil {
		return nil, err
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

func (s *orderServiceImpl) FindByID(id uint64) (*modelRes.Order, error) {

	order, err := s.orderRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return s.orderEntityToModelRes(order), nil
}

func (s *orderServiceImpl) Update(id uint64, order *modelReq.Order) (*modelRes.Order, error) {

	orderEntity := s.orderReqToEntity(order)
	orderEntity, err := s.orderRepository.Update(id, orderEntity)
	if err != nil {
		return nil, err
	}
	return s.orderEntityToModelRes(orderEntity), nil
}

func (s *orderServiceImpl) Delete(id uint64) error {

	err := s.orderRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (s *orderServiceImpl) ChangeStatusNext(id uint64) (*modelRes.Order, error) {

	order, err := s.orderRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	err = order.NextStatus()
	if err != nil {
		return nil, err
	}

	_, err = s.orderRepository.Update(id, order)
	if err != nil {
		return nil, err
	}

	return s.orderEntityToModelRes(order), nil
}
func (s *orderServiceImpl) ChageStatusDone(id uint64) (*modelRes.Order, error) {

	order, err := s.orderRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	err = order.NextPaidToDone()
	if err != nil {
		return nil, err
	}

	_, err = s.orderRepository.Update(id, order)
	if err != nil {
		return nil, err
	}

	return s.orderEntityToModelRes(order), nil
}

func (s *orderServiceImpl) orderReqToEntity(orderReq *modelReq.Order) *entities.Order {
	return &entities.Order{
		TransactionID: orderReq.TransactionID,
		ProductID:     orderReq.ProductID,
		Status:        orderReq.Status,
	}
}

func (s *orderServiceImpl) orderEntityToModelRes(order *entities.Order) *modelRes.Order {
	return &modelRes.Order{
		OrderID:       order.OrderID,
		TransactionID: order.TransactionID,
		ProductID:     order.ProductID,
		Status:        order.Status,
	}
}
