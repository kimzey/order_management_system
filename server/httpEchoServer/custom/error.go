package custom

import (
	"errors"
	"github.com/labstack/echo/v4"
)

type ErrorMessage struct {
	ErrorCode int    `json:"StatusCode"`
	Message   string `json:"Error"`
}

func Error(c echo.Context, statusCode int, err error) error {
	return c.JSON(
		statusCode,
		&ErrorMessage{Message: err.Error(), ErrorCode: statusCode},
	)
}

var (
	ErrInvalidRequest               = errors.New("1001: Invalid request data")
	ErrOrderNotFound                = errors.New("1002: Order not found")
	ErrFailedToCreateOrder          = errors.New("1003: Failed to create order")
	ErrFailedToUpdateOrder          = errors.New("1004: Failed to update order")
	ErrFailedToDeleteOrder          = errors.New("1005: Failed to delete order")
	ErrFailedToRetrieveOrders       = errors.New("1006: Failed to retrieve orders")
	ErrFailedToChangeOrderStatus    = errors.New("1007: Failed to change order status")
	ErrFailedToValidateOrderRequest = errors.New("1008: Failed to validate order request")

	ErrFailedToCreateProduct          = errors.New("2001: Failed to create product")
	ErrFailedToUpdateProduct          = errors.New("2002: Failed to update product")
	ErrFailedToDeleteProduct          = errors.New("2003: Failed to delete product")
	ErrFailedToRetrieveProducts       = errors.New("2004: Failed to retrieve products")
	ErrProductNotFound                = errors.New("2005: Product not found")
	ErrFailedToValidateProductRequest = errors.New("2006: Failed to validate product request")

	ErrFailedToCreateStock          = errors.New("3001: Failed to create stock")
	ErrFailedToUpdateStock          = errors.New("3002: Failed to update stock")
	ErrFailedToDeleteStock          = errors.New("3003: Failed to delete stock")
	ErrFailedToRetrieveStocks       = errors.New("3004: Failed to retrieve stocks")
	ErrStockNotFound                = errors.New("3005: Stock not found")
	ErrFailedToValidateStockRequest = errors.New("3006: Failed to validate stock request")

	ErrFailedToCreateTransaction          = errors.New("4001: Failed to create transaction")
	ErrFailedToUpdateTransaction          = errors.New("4002: Failed to update transaction")
	ErrFailedToDeleteTransaction          = errors.New("4003: Failed to delete transaction")
	ErrFailedToRetrieveTransactions       = errors.New("4004: Failed to retrieve transactions")
	ErrTransactionNotFound                = errors.New("4005: Transaction not found")
	ErrFailedToValidateTransactionRequest = errors.New("4006: Failed to validate transaction request")
)
