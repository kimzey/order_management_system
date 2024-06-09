package main

import (
	"github.com/kizmey/order_management_system/config"
	"github.com/kizmey/order_management_system/database"
	"github.com/kizmey/order_management_system/pkg"
	serverPkg "github.com/kizmey/order_management_system/server/httpEchoServer"

	//_orderController "github.com/kizmey/order_management_system/pkg/order/controller"
	_orderRepository "github.com/kizmey/order_management_system/pkg/order/repository"
	_orderService "github.com/kizmey/order_management_system/pkg/order/service"

	//_productController "github.com/kizmey/order_management_system/pkg/product/controller"
	_productRepository "github.com/kizmey/order_management_system/pkg/product/repository"
	_productService "github.com/kizmey/order_management_system/pkg/product/service"

	//_stockController "github.com/kizmey/order_management_system/pkg/stock/controller"
	_stockRepository "github.com/kizmey/order_management_system/pkg/stock/repository"
	_stockService "github.com/kizmey/order_management_system/pkg/stock/service"

	//_transactionController "github.com/kizmey/order_management_system/pkg/transaction/controller"
	_transactionRepository "github.com/kizmey/order_management_system/pkg/transaction/repository"
	_transactionService "github.com/kizmey/order_management_system/pkg/transaction/service"
)

func main() {
	conf := config.GettingConfig()
	db := database.NewPostgresDatabase(conf.Database)

	usecase := initUsecase(db)

	server := serverPkg.NewEchoServer(conf, usecase)
	server.Start()
}

func initUsecase(db database.Database) *pkg.Usecase {

	// Init Repository
	orderRepo := _orderRepository.NewOrderRepositoryImpl(db)
	productRepo := _productRepository.NewProductRepositoryImpl(db)
	stockRepo := _stockRepository.NewStockRepositoryImpl(db)
	transactionRepo := _transactionRepository.NewTransactionController(db)

	// Init Service
	productService := _productService.NewProductServiceImpl(productRepo)
	stockService := _stockService.NewStockServiceImpl(stockRepo)
	transactionService := _transactionService.NewTransactionServiceImpl(transactionRepo)
	orderService := _orderService.NewOrderServiceImpl(orderRepo, transactionRepo, stockRepo, productRepo)

	usecase := pkg.NewUsecase(transactionService, stockService, productService, orderService)

	return usecase
}
