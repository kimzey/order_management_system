package httpEchoServer

import (
	_orderController "github.com/kizmey/order_management_system/pkg/controller"
)

func (s *echoServer) initOrderRouter() {
	router := s.app.Group("/v1/order")

	orderController := _orderController.NewOrderControllerImpl(s.usecase.OrderService)

	router.POST("", orderController.Create)
	router.GET("", orderController.FindAll)
	router.GET("/:id", orderController.FindByID)
	router.PUT("/:id", orderController.Update)
	router.DELETE("/:id", orderController.Delete)
	router.PUT("/next/:id", orderController.ChangeStatusNext)
	router.PUT("/done/:id", orderController.ChageStatusDone)
}
