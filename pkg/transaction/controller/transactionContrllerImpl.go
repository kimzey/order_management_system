package controller

import (
	"github.com/kizmey/order_management_system/entities"
	_transactionService "github.com/kizmey/order_management_system/pkg/transaction/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type transactionControllerImpl struct {
	transaction _transactionService.TransactionService
}

func NewTransactionController(transaction _transactionService.TransactionService) TransactionController {
	return &transactionControllerImpl{
		transaction: transaction,
	}
}

func (c *transactionControllerImpl) Create(ctx echo.Context) error {
	transactionReq := new(entities.Transaction)
	if err := ctx.Bind(transactionReq); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	transaction, err := c.transaction.Create(transactionReq)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusCreated, transaction)
}

func (c *transactionControllerImpl) FindAll(ctx echo.Context) error {
	transactions, err := c.transaction.FindAll()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, transactions)
}
