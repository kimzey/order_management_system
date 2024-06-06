package server

import (
	_stockController "github.com/kizmey/order_management_system/pkg/stock/controller"
	_stockRepository "github.com/kizmey/order_management_system/pkg/stock/repository"
	_stockService "github.com/kizmey/order_management_system/pkg/stock/service"
)

func (s *echoServer) initStockRouter() {
	router := s.app.Group("/v1/stock")

	stockRepository := _stockRepository.NewStockRepositoryImpl(s.db, s.app.Logger)
	stockService := _stockService.NewStockServiceImpl(stockRepository)
	stockController := _stockController.NewStockControllerImpl(stockService)

	router.GET("", stockController.FindAll)
	router.GET("/product/:id", stockController.CheckStockByProductId)
	router.POST("", stockController.Create)
	router.PUT("/:id", stockController.Update)
	router.DELETE("/:id", stockController.Delete)
}
