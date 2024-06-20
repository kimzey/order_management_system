package httpEchoServer

import (
	_transactionController "github.com/kizmey/order_management_system/pkg/controller/transaction"
)

func (s *echoServer) inittransactionRouter() {
	router := s.app.Group("/v1/transaction")

	transactionController := _transactionController.NewTransactionControllerImpl(s.usecase.TransactionService)

	router.POST("", transactionController.Create)
	router.GET("", transactionController.FindAll)
	router.GET("/:id", transactionController.FindByID)
	router.PUT("/:id", transactionController.Update)
	router.DELETE("/:id", transactionController.Delete)

}
