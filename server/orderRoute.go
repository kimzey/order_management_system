package server

import (
	_orderController "github.com/kizmey/order_management_system/pkg/order/controller"
	_orderRepository "github.com/kizmey/order_management_system/pkg/order/repository"
	_orderService "github.com/kizmey/order_management_system/pkg/order/service"
	_productRepository "github.com/kizmey/order_management_system/pkg/product/repository"
	_stockRepository "github.com/kizmey/order_management_system/pkg/stock/repository"
	_transactionRepository "github.com/kizmey/order_management_system/pkg/transaction/repository"
)

func (s *echoServer) initOrderRouter() {
	router := s.app.Group("/v1/order")

	productRepository := _productRepository.NewProductRepositoryImpl(s.db, s.app.Logger)
	transactionRepository := _transactionRepository.NewTransactionController(s.db, s.app.Logger)
	orderRepository := _orderRepository.NewOrderRepositoryImpl(s.db, s.app.Logger)
	stockRepository := _stockRepository.NewStockRepositoryImpl(s.db, s.app.Logger)

	orderService := _orderService.NewOrderService(orderRepository, transactionRepository, stockRepository, productRepository)
	orderController := _orderController.NewOrderControllerImpl(orderService)

	router.POST("", orderController.Create)
	router.GET("", orderController.FindAll)
	router.GET("/:id", orderController.FindByID)
	router.PUT("/:id", orderController.Update)
	router.DELETE("/:id", orderController.Delete)
	router.PUT("/next/:id", orderController.ChangeStatusNext)
	router.PUT("/done/:id", orderController.ChageStatusDone)

}
