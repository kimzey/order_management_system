package transaction

import (
	"github.com/kizmey/order_management_system/pkg/interface/modelReq"
	_transactionService "github.com/kizmey/order_management_system/pkg/service/transaction"
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
	ctx, sp := tracer.Start(pctx.Request().Context(), "transactionCreateController")
	defer sp.End()

	transactionReq := new(modelReq.Transaction)

	validatingContext := custom.NewCustomEchoRequest(pctx)
	if err := validatingContext.BindAndValidate(transactionReq); err != nil {
		return custom.Error(pctx, http.StatusBadRequest, custom.ErrFailedToValidateTransactionRequest)
	}

	if !isProductIDsUnique(transactionReq.Product) {
		return custom.Error(pctx, http.StatusBadRequest, custom.ErrFailedToValidateTransactionRequest)
	}

	transaction, err := c.transaction.Create(ctx, transactionReq)
	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, custom.ErrFailedToCreateTransaction)
	}

	return pctx.JSON(http.StatusCreated, transaction)
}

func (c *transactionControllerImpl) FindAll(pctx echo.Context) error {
	ctx, sp := tracer.Start(pctx.Request().Context(), "transactionFindAllController")
	defer sp.End()

	transactions, err := c.transaction.FindAll(ctx)
	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, custom.ErrFailedToRetrieveTransactions)
	}
	return pctx.JSON(http.StatusOK, transactions)
}

func (c *transactionControllerImpl) FindByID(pctx echo.Context) error {
	ctx, sp := tracer.Start(pctx.Request().Context(), "transactionFindByIdController")
	defer sp.End()

	id := pctx.Param("id")

	transaction, err := c.transaction.FindByID(ctx, id)
	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, custom.ErrTransactionNotFound)
	}

	return pctx.JSON(http.StatusOK, transaction)
}

func (c *transactionControllerImpl) Update(pctx echo.Context) error {
	ctx, sp := tracer.Start(pctx.Request().Context(), "transactionUpdateController")
	defer sp.End()

	id := pctx.Param("id")
	transactionReq := new(modelReq.Transaction)

	validatingContext := custom.NewCustomEchoRequest(pctx)
	if err := validatingContext.BindAndValidate(transactionReq); err != nil {
		return custom.Error(pctx, http.StatusBadRequest, custom.ErrFailedToValidateTransactionRequest)
	}

	if !isProductIDsUnique(transactionReq.Product) {
		return custom.Error(pctx, http.StatusBadRequest, custom.ErrFailedToValidateTransactionRequest)
	}

	transaction, err := c.transaction.Update(ctx, id, transactionReq)
	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, custom.ErrFailedToUpdateTransaction)
	}

	return pctx.JSON(http.StatusOK, transaction)
}

func (c *transactionControllerImpl) Delete(pctx echo.Context) error {
	ctx, sp := tracer.Start(pctx.Request().Context(), "transactionDeleteController")
	defer sp.End()

	id := pctx.Param("id")

	transaction, err := c.transaction.Delete(ctx, id)
	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, custom.ErrFailedToDeleteTransaction)
	}

	return pctx.JSON(http.StatusOK, transaction)
}

func isProductIDsUnique(products []modelReq.ProductItem) bool {
	seen := make(map[string]struct{})
	for _, item := range products {
		if _, found := seen[item.ProductID]; found {
			return false
		}
		seen[item.ProductID] = struct{}{}
	}
	return true
}
