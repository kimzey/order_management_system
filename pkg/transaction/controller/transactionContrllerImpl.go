package controller

import (
	"github.com/kizmey/order_management_system/pkg/modelReq"
	_transactionService "github.com/kizmey/order_management_system/pkg/transaction/service"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type transactionControllerImpl struct {
	transaction _transactionService.TransactionService
}

func NewTransactionControllerImpl(transaction _transactionService.TransactionService) TransactionController {
	return &transactionControllerImpl{
		transaction: transaction,
	}
}

func (c *transactionControllerImpl) Create(pctx echo.Context) error {
	transactionReq := new(modelReq.Transaction)
	if err := pctx.Bind(transactionReq); err != nil {
		return pctx.JSON(http.StatusBadRequest, err.Error())
	}

	transaction, err := c.transaction.Create(transactionReq)
	if err != nil {
		return pctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return pctx.JSON(http.StatusCreated, transaction)
}

func (c *transactionControllerImpl) FindAll(pctx echo.Context) error {
	transactions, err := c.transaction.FindAll()
	if err != nil {
		return pctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return pctx.JSON(http.StatusOK, transactions)
}

func (c *transactionControllerImpl) FindByID(pctx echo.Context) error {
	id, err := strconv.ParseUint(pctx.Param("id"), 0, 64)
	if err != nil {
		return pctx.JSON(http.StatusBadRequest, err.Error())
	}

	transaction, err := c.transaction.FindByID(id)
	if err != nil {
		return pctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return pctx.JSON(http.StatusOK, transaction)
}

func (c *transactionControllerImpl) Update(pctx echo.Context) error {

	id, err := strconv.ParseUint(pctx.Param("id"), 0, 64)
	if err != nil {
		return pctx.JSON(http.StatusBadRequest, err.Error())
	}

	transactionReq := new(modelReq.Transaction)
	if err := pctx.Bind(transactionReq); err != nil {
		return pctx.JSON(http.StatusBadRequest, err.Error())
	}

	transaction, err := c.transaction.Update(id, transactionReq)
	if err != nil {
		return pctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return pctx.JSON(http.StatusOK, transaction)
}

func (c *transactionControllerImpl) Delete(pctx echo.Context) error {
	id, err := strconv.ParseUint(pctx.Param("id"), 0, 64)
	if err != nil {
		return pctx.JSON(http.StatusBadRequest, err.Error())
	}

	err = c.transaction.Delete(id)
	if err != nil {
		return pctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return pctx.JSON(http.StatusOK, "deleted")
}
