package server

import (
	_orderController "github.com/kizmey/order_management_system/pkg/order/controller"
	_orderRepository "github.com/kizmey/order_management_system/pkg/order/repository"
	_orderService "github.com/kizmey/order_management_system/pkg/order/service"
	_transactionRepository "github.com/kizmey/order_management_system/pkg/transaction/repository"
)

func (s *echoServer) initOrderRouter() {
	router := s.app.Group("/v1/order")

	transactionRepository := _transactionRepository.NewTransactionController(s.db, s.app.Logger)
	orderRepository := _orderRepository.NewOrderRepositoryImpl(s.db, s.app.Logger)

	orderService := _orderService.NewOrderService(orderRepository, transactionRepository)
	orderController := _orderController.NewOrderControllerImpl(orderService)

	router.GET("", orderController.FindAll)
	//router.GET("/product/:id", stockController.CheckStockByProductId)
	router.POST("", orderController.Create)
	router.PUT("/next/:id", orderController.ChangeStatusNext)
	router.PUT("/done/:id", orderController.ChageStatusDone)

	//router.PUT("/:id", stockController.Update)
	//router.DELETE("/:id", stockController.Delete)
}
