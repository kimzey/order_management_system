package httpEchoServer

import (
	_productController "github.com/kizmey/order_management_system/pkg/product/controller"
)

func (s *echoServer) initProductRouter() {
	router := s.app.Group("/v1/product")

	productController := _productController.NewProductControllerImpl(s.usecase.ProductService)

	router.POST("", productController.Create)
	router.GET("", productController.FindAll)
	router.GET("/:id", productController.FindByID)
	router.PUT("/:id", productController.Update)
	router.DELETE("/:id", productController.Delete)
}
