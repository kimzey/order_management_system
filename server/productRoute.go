package server

import (
	_productController "github.com/kizmey/order_management_system/pkg/product/controller"
	_productRepository "github.com/kizmey/order_management_system/pkg/product/repository"
	_productService "github.com/kizmey/order_management_system/pkg/product/service"
)

func (s *echoServer) initProductRouter() {
	router := s.app.Group("/v1/product")

	productRepository := _productRepository.NewProductRepositoryImpl(s.db, s.app.Logger)
	productService := _productService.NewProductServiceImpl(productRepository)
	productController := _productController.NewProductController(productService)

	router.POST("", productController.Create)
	router.GET("", productController.FindAll)
	router.GET("/:id", productController.FindByID)
	router.PUT("/:id", productController.Update)
	router.DELETE("/:id", productController.Delete)
}
