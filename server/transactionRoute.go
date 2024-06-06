package server

import (
	_transactionController "github.com/kizmey/order_management_system/pkg/transaction/controller"
	_transactionRepository "github.com/kizmey/order_management_system/pkg/transaction/repository"
	_transactionService "github.com/kizmey/order_management_system/pkg/transaction/service"
)

func (s *echoServer) inittransactionRouter() {
	router := s.app.Group("/v1/transaction")

	transactionRepository := _transactionRepository.NewTransactionController(s.db, s.app.Logger)
	transactionService := _transactionService.NewTransactionService(transactionRepository)
	transactionController := _transactionController.NewTransactionController(transactionService)

	router.POST("", transactionController.Create)
	router.GET("", transactionController.FindAll)

}
