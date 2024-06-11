package controller

import (
	"fmt"
	"github.com/kizmey/order_management_system/pkg/interface/modelReq"
	_transactionService "github.com/kizmey/order_management_system/pkg/service"
	"github.com/kizmey/order_management_system/server/httpEchoServer/custom"
	"github.com/labstack/echo/v4"
	"net/http"
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

	validatingContext := custom.NewCustomEchoRequest(pctx)
	if err := validatingContext.BindAndValidate(transactionReq); err != nil {
		return custom.Error(pctx, http.StatusBadRequest, err)
	}

	transaction, err := c.transaction.Create(transactionReq)
	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, err)
	}

	return pctx.JSON(http.StatusCreated, transaction)
}

func (c *transactionControllerImpl) FindAll(pctx echo.Context) error {
	transactions, err := c.transaction.FindAll()
	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, err)
	}
	return pctx.JSON(http.StatusOK, transactions)
}

func (c *transactionControllerImpl) FindByID(pctx echo.Context) error {
	id := pctx.Param("id")
	fmt.Println(id)
	//id, err := custom.CheckParamId(pctx)
	//if err != nil {
	//	return custom.Error(pctx, http.StatusBadRequest, err)
	//}

	transaction, err := c.transaction.FindByID(id)
	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, err)
	}

	return pctx.JSON(http.StatusOK, transaction)
}

func (c *transactionControllerImpl) Update(pctx echo.Context) error {
	id := pctx.Param("id")
	fmt.Println(id)
	//id, err := custom.CheckParamId(pctx)
	//if err != nil {
	//	return custom.Error(pctx, http.StatusBadRequest, err)
	//}

	transactionReq := new(modelReq.Transaction)

	validatingContext := custom.NewCustomEchoRequest(pctx)
	if err := validatingContext.BindAndValidate(transactionReq); err != nil {
		return custom.Error(pctx, http.StatusBadRequest, err)
	}

	transaction, err := c.transaction.Update(id, transactionReq)
	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, err)
	}

	return pctx.JSON(http.StatusOK, transaction)
}

func (c *transactionControllerImpl) Delete(pctx echo.Context) error {
	id := pctx.Param("id")
	fmt.Println(id)
	//id, err := custom.CheckParamId(pctx)
	//if err != nil {
	//	return custom.Error(pctx, http.StatusBadRequest, err)
	//}

	err := c.transaction.Delete(id)
	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, err)
	}

	return pctx.JSON(http.StatusOK, "deleted successfully")
}
