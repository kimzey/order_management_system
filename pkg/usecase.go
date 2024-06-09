package pkg

import (
	_orderService "github.com/kizmey/order_management_system/pkg/order/service"
	_productService "github.com/kizmey/order_management_system/pkg/product/service"
	_stockService "github.com/kizmey/order_management_system/pkg/stock/service"
	_transactionService "github.com/kizmey/order_management_system/pkg/transaction/service"
)

type Usecase struct {
	TransactionService _transactionService.TransactionService
	StockService       _stockService.StockService
	ProductService     _productService.ProductService
	OrderService       _orderService.OrderService
}

func NewUsecase(
	transactionService _transactionService.TransactionService,
	stockService _stockService.StockService,
	productService _productService.ProductService,
	orderService _orderService.OrderService,
) *Usecase {
	return &Usecase{
		TransactionService: transactionService,
		StockService:       stockService,
		ProductService:     productService,
		OrderService:       orderService,
	}
}
