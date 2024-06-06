package service

import (
	"github.com/kizmey/order_management_system/entities"
	_orderRepository "github.com/kizmey/order_management_system/pkg/order/repository"
	_transaction "github.com/kizmey/order_management_system/pkg/transaction/repository"
)

type orderServiceImpl struct {
	orderRepository       _orderRepository.OrderRepository
	transactionRepository _transaction.TransactionRepository
}

func NewOrderService(orderRepository _orderRepository.OrderRepository, transaction _transaction.TransactionRepository) OrderService {
	return &orderServiceImpl{
		orderRepository:       orderRepository,
		transactionRepository: transaction,
	}
}

func (s *orderServiceImpl) Create(order *entities.Order) (*entities.Order, error) {

	newTransaction := &entities.Transaction{
		ProductName:  order.ProductName,
		ProductPrice: order.ProductPrice,
		Quantity:     order.Quantity,
		SumPrice:     s.calculatePrice(order.ProductPrice, order.Quantity, order.IsDomestic),
	}

	tranID, err := s.transactionRepository.Create(newTransaction)

	order.TransactionID = tranID
	newOrder, err := s.orderRepository.Create(order)

	if err != nil {
		return nil, err
	}

	return newOrder, nil
}

func (s *orderServiceImpl) ChangeStatusNext(id uint64) (*entities.Order, error) {
	return s.orderRepository.ChangeStatusNext(id)
}
func (s *orderServiceImpl) ChageStatusDone(id uint64) (*entities.Order, error) {
	return s.orderRepository.ChageStatusDone(id)
}

func (s *orderServiceImpl) calculatePrice(price uint, quantity uint, isDomestic bool) uint {
	if isDomestic {
		return (price * quantity) + 40
	} else {
		return (price * quantity) + 200
	}
}
