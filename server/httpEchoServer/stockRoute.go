package httpEchoServer

import (
	_stockController "github.com/kizmey/order_management_system/pkg/stock/controller"
)

func (s *echoServer) initStockRouter() {
	router := s.app.Group("/v1/stock")

	stockController := _stockController.NewStockControllerImpl(s.usecase.StockService)

	router.GET("", stockController.FindAll)
	router.GET("/product/:id", stockController.CheckStockByProductId)
	router.POST("", stockController.Create)
	router.PUT("/:id", stockController.Update)
	router.DELETE("/:id", stockController.Delete)
}
