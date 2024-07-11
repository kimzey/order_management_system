package transaction

import (
	"errors"
	_interface "github.com/kizmey/order_management_system/pkg/interface/aggregation"
	"github.com/kizmey/order_management_system/pkg/interface/entities"
	"github.com/kizmey/order_management_system/pkg/interface/modelReq"
	"github.com/kizmey/order_management_system/pkg/interface/modelRes"
	_transactionService "github.com/kizmey/order_management_system/pkg/service/transaction"
	"github.com/kizmey/order_management_system/server/httpEchoServer/custom"
	"github.com/labstack/echo/v4"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
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

	transactionEntity := c.transactionReqToAggregation(transactionReq)

	transaction, err := c.transaction.Create(ctx, transactionEntity)
	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, custom.ErrFailedToCreateTransaction)
	}

	transactionRes := c.transactionAndproductEntityToRes(transaction)
	c.SetSubAttributesWithJson(transactionRes, sp)

	return pctx.JSON(http.StatusCreated, transactionRes)
}

func (c *transactionControllerImpl) FindAll(pctx echo.Context) error {
	ctx, sp := tracer.Start(pctx.Request().Context(), "transactionFindAllController")
	defer sp.End()

	transactions, err := c.transaction.FindAll(ctx)
	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, custom.ErrFailedToRetrieveTransactions)
	}

	allTransaction := make([]modelRes.Transaction, 0)
	for _, transactionEntity := range *transactions {
		allTransaction = append(allTransaction, *c.transactionEntityToRes(&transactionEntity))
	}
	c.SetSubAttributesWithJson(allTransaction, sp)

	return pctx.JSON(http.StatusOK, allTransaction)
}

func (c *transactionControllerImpl) FindByID(pctx echo.Context) error {
	ctx, sp := tracer.Start(pctx.Request().Context(), "transactionFindByIdController")
	defer sp.End()

	id := pctx.Param("id")

	transaction, err := c.transaction.FindByID(ctx, id)
	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, custom.ErrTransactionNotFound)
	}

	transactionRes := c.transactionEntityToRes(transaction)
	c.SetSubAttributesWithJson(transactionRes, sp)

	return pctx.JSON(http.StatusOK, transactionRes)
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

	transaction := c.transactionReqToAggregation(transactionReq)
	transaction, err := c.transaction.Update(ctx, id, transaction)
	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, custom.ErrFailedToUpdateTransaction)
	}

	transactionRes := c.transactionAndproductEntityToRes(transaction)
	c.SetSubAttributesWithJson(transactionRes, sp)

	return pctx.JSON(http.StatusOK, transactionRes)
}

func (c *transactionControllerImpl) Delete(pctx echo.Context) error {
	ctx, sp := tracer.Start(pctx.Request().Context(), "transactionDeleteController")
	defer sp.End()

	id := pctx.Param("id")

	transaction, err := c.transaction.Delete(ctx, id)
	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, custom.ErrFailedToDeleteTransaction)
	}

	transactionRes := c.transactionEntityToRes(transaction)
	c.SetSubAttributesWithJson(transactionRes, sp)

	return pctx.JSON(http.StatusOK, transactionRes)
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

func (c *transactionControllerImpl) transactionReqToAggregation(transactionReq *modelReq.Transaction) *_interface.TransactionEcommerce {
	mapProduct := make(map[string]uint)
	for _, item := range transactionReq.Product {
		mapProduct[item.ProductID] = item.Quantity
	}

	transaction := entities.Transaction{
		IsDomestic: transactionReq.IsDomestic,
	}

	return _interface.NewTransactionEcommerce(&transaction, nil, mapProduct)
}

func (c *transactionControllerImpl) transactionAndproductEntityToRes(transactionEntity *_interface.TransactionEcommerce) *modelRes.Transaction {
	products := make([]modelRes.Product, 0)

	for _, product := range transactionEntity.Product {
		products = append(products, modelRes.Product{
			ProductID:   product.ProductID,
			ProductName: product.ProductName,
			Quantity:    transactionEntity.AddessProduct[product.ProductID],
		})
	}

	return &modelRes.Transaction{
		TransactionID: transactionEntity.Tranasaction.TransactionID,
		IsDomestic:    transactionEntity.Tranasaction.IsDomestic,
		SumPrice:      transactionEntity.Tranasaction.SumPrice,
		Products:      products,
	}
}

func (c *transactionControllerImpl) transactionEntityToRes(transactionEntity *entities.Transaction) *modelRes.Transaction {
	return &modelRes.Transaction{
		TransactionID: transactionEntity.TransactionID,
		IsDomestic:    transactionEntity.IsDomestic,
		SumPrice:      transactionEntity.SumPrice,
	}

}

func (c *transactionControllerImpl) SetSubAttributesWithJson(obj any, sp trace.Span) {
	if transactions, ok := obj.([]modelRes.Transaction); ok {
		var transactionIDs []string
		var transactionIsDomestic []bool
		var transactionSumPrices []int
		var transactionProducts []string

		for _, transaction := range transactions {
			transactionIDs = append(transactionIDs, transaction.TransactionID)
			transactionIsDomestic = append(transactionIsDomestic, transaction.IsDomestic)
			transactionSumPrices = append(transactionSumPrices, int(transaction.SumPrice))
			for _, product := range transaction.Products {
				transactionProducts = append(transactionProducts, product.ProductID)
			}
		}

		sp.SetAttributes(
			attribute.StringSlice("TransactionID", transactionIDs),
			attribute.BoolSlice("TransactionIsDomestic", transactionIsDomestic),
			attribute.IntSlice("TransactionSumPrice", transactionSumPrices),
			// attribute.StringSlice("TransactionProducts", transactionProducts),
		)
	} else if transaction, ok := obj.(*modelRes.Transaction); ok {
		var transactionProducts []string
		for _, product := range transaction.Products {
			transactionProducts = append(transactionProducts, product.ProductID)
		}

		sp.SetAttributes(
			attribute.String("TransactionID", transaction.TransactionID),
			attribute.Bool("TransactionIsDomestic", transaction.IsDomestic),
			attribute.Int("TransactionSumPrice", int(transaction.SumPrice)),
			// attribute.StringSlice("TransactionProducts", transactionProducts),
		)
	} else {
		sp.RecordError(errors.New("invalid type"))
	}
}
