package service

import (
	"fmt"
	"github.com/kizmey/order_management_system/entities"
	_orderRepository "github.com/kizmey/order_management_system/pkg/order/repository"
	_productRepository "github.com/kizmey/order_management_system/pkg/product/repository"
	_stockRepository "github.com/kizmey/order_management_system/pkg/stock/repository"
	_transactionRepository "github.com/kizmey/order_management_system/pkg/transaction/repository"
)

type orderServiceImpl struct {
	orderRepository       _orderRepository.OrderRepository
	transactionRepository _transactionRepository.TransactionRepository
	stockRepository       _stockRepository.StockRepository
	productRepository     _productRepository.ProductRepository
}

func NewOrderService(orderRepository _orderRepository.OrderRepository,
	transactionRepository _transactionRepository.TransactionRepository,
	stockRepository _stockRepository.StockRepository,
	productRepository _productRepository.ProductRepository,
) OrderService {
	return &orderServiceImpl{
		orderRepository:       orderRepository,
		transactionRepository: transactionRepository,
		stockRepository:       stockRepository,
		productRepository:     productRepository,
	}
}

func (s *orderServiceImpl) Create(order *entities.Order) (*entities.Order, error) {

	stock, err := s.stockRepository.CheckStockByProductId(order.ProductID)
	if err != nil {
		return nil, err
	}
	if stock.Quantity < order.Quantity {
		return nil, fmt.Errorf("stock not enough")
	}

	stock.Quantity = stock.Quantity - order.Quantity
	_, err = s.stockRepository.Update(stock.StockID, stock)
	if err != nil {
		return nil, err
	}

	checkProduct, err := s.productRepository.FindByID(order.ProductID)
	if err != nil {
		return nil, err
	}

	newTransaction := &entities.Transaction{
		ProductID:    order.ProductID,
		ProductName:  checkProduct.ProductName,
		ProductPrice: checkProduct.ProductPrice,
		Quantity:     order.Quantity,
		SumPrice:     s.calculatePrice(checkProduct.ProductPrice, order.Quantity, order.IsDomestic),
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
	return s.orderRepository.ChangeStatusDone(id)
}

func (s *orderServiceImpl) FindAll() (*[]entities.Order, error) {
	return s.orderRepository.FindAll()
}

func (s *orderServiceImpl) calculatePrice(price uint, quantity uint, isDomestic bool) uint {
	if isDomestic {
		return (price * quantity) + 40
	} else {
		return (price * quantity) + 200
	}
}
