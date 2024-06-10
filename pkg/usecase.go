package pkg

import (
	"github.com/kizmey/order_management_system/database"
	_orderRepository "github.com/kizmey/order_management_system/pkg/order/repository"
	_orderService "github.com/kizmey/order_management_system/pkg/order/service"
	_productRepository "github.com/kizmey/order_management_system/pkg/product/repository"
	_productService "github.com/kizmey/order_management_system/pkg/product/service"
	_stockRepository "github.com/kizmey/order_management_system/pkg/stock/repository"
	_stockService "github.com/kizmey/order_management_system/pkg/stock/service"
	_transactionRepository "github.com/kizmey/order_management_system/pkg/transaction/repository"
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

func InitUsecase(db database.Database) *Usecase {

	// Init Repository
	orderRepo := _orderRepository.NewOrderRepositoryImpl(db)
	productRepo := _productRepository.NewProductRepositoryImpl(db)
	stockRepo := _stockRepository.NewStockRepositoryImpl(db)
	transactionRepo := _transactionRepository.NewTransactionController(db)

	// Init Service
	productService := _productService.NewProductServiceImpl(productRepo)
	stockService := _stockService.NewStockServiceImpl(stockRepo)
	transactionService := _transactionService.NewTransactionServiceImpl(transactionRepo, stockRepo, productRepo)
	orderService := _orderService.NewOrderServiceImpl(orderRepo, transactionRepo, stockRepo, productRepo)

	usecase := NewUsecase(transactionService, stockService, productService, orderService)

	return usecase
}
