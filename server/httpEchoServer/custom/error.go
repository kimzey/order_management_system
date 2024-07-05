package custom

import (
	"github.com/labstack/echo/v4"
)

type ErrorMessage struct {
	ErrorCode int    `json:"ErrorCode"`
	Message   string `json:"Error"`
}

func Error(c echo.Context, statusCode int, err ErrorStatus) error {
	return c.JSON(
		statusCode,
		ErrorMessage{
			ErrorCode: err.Code,
			Message:   err.Message,
		},
	)
}

type ErrorStatus struct {
	Code    int
	Message string
}

var (
	//ErrInvalidRequest = ErrorStatus{Code: 1001, Message: "Invalid request data"}

	ErrOrderNotFound                = ErrorStatus{Code: 1002, Message: "Order not found"}
	ErrFailedToCreateOrder          = ErrorStatus{Code: 1003, Message: "Failed to create order"}
	ErrFailedToUpdateOrder          = ErrorStatus{Code: 1004, Message: "Failed to update order"}
	ErrFailedToDeleteOrder          = ErrorStatus{Code: 1005, Message: "Failed to delete order"}
	ErrFailedToRetrieveOrders       = ErrorStatus{Code: 1006, Message: "Failed to retrieve orders"}
	ErrFailedToChangeOrderStatus    = ErrorStatus{Code: 1007, Message: "Failed to change order status"}
	ErrFailedToValidateOrderRequest = ErrorStatus{Code: 1008, Message: "Failed to validate order request"}

	ErrFailedToCreateProduct          = ErrorStatus{Code: 2001, Message: "Failed to create product"}
	ErrFailedToUpdateProduct          = ErrorStatus{Code: 2002, Message: "Failed to update product"}
	ErrFailedToDeleteProduct          = ErrorStatus{Code: 2003, Message: "Failed to delete product"}
	ErrFailedToRetrieveProducts       = ErrorStatus{Code: 2004, Message: "Failed to retrieve products"}
	ErrProductNotFound                = ErrorStatus{Code: 2005, Message: "Product not found"}
	ErrFailedToValidateProductRequest = ErrorStatus{Code: 2006, Message: "Failed to validate product request"}

	ErrFailedToCreateStock          = ErrorStatus{Code: 3001, Message: "Failed to create stock"}
	ErrFailedToUpdateStock          = ErrorStatus{Code: 3002, Message: "Failed to update stock"}
	ErrFailedToDeleteStock          = ErrorStatus{Code: 3003, Message: "Failed to delete stock"}
	ErrFailedToRetrieveStocks       = ErrorStatus{Code: 3004, Message: "Failed to retrieve stocks"}
	ErrStockNotFound                = ErrorStatus{Code: 3005, Message: "Stock not found"}
	ErrFailedToValidateStockRequest = ErrorStatus{Code: 3006, Message: "Failed to validate stock request"}

	ErrFailedToCreateTransaction          = ErrorStatus{Code: 4001, Message: "Failed to create transaction"}
	ErrFailedToUpdateTransaction          = ErrorStatus{Code: 4002, Message: "Failed to update transaction"}
	ErrFailedToDeleteTransaction          = ErrorStatus{Code: 4003, Message: "Failed to delete transaction"}
	ErrFailedToRetrieveTransactions       = ErrorStatus{Code: 4004, Message: "Failed to retrieve transactions"}
	ErrTransactionNotFound                = ErrorStatus{Code: 4005, Message: "Transaction not found"}
	ErrFailedToValidateTransactionRequest = ErrorStatus{Code: 4006, Message: "Failed to validate transaction request"}
)

//var (
//	ErrInvalidRequest               = errors.New("1001: Invalid request data")
//	ErrOrderNotFound                = errors.New("1002: Order not found")
//	ErrFailedToCreateOrder          = errors.New("1003: Failed to create order")
//	ErrFailedToUpdateOrder          = errors.New("1004: Failed to update order")
//	ErrFailedToDeleteOrder          = errors.New("1005: Failed to delete order")
//	ErrFailedToRetrieveOrders       = errors.New("1006: Failed to retrieve orders")
//	ErrFailedToChangeOrderStatus    = errors.New("1007: Failed to change order status")
//	ErrFailedToValidateOrderRequest = errors.New("1008: Failed to validate order request")
//
//	ErrFailedToCreateProduct          = errors.New("2001: Failed to create product")
//	ErrFailedToUpdateProduct          = errors.New("2002: Failed to update product")
//	ErrFailedToDeleteProduct          = errors.New("2003: Failed to delete product")
//	ErrFailedToRetrieveProducts       = errors.New("2004: Failed to retrieve products")
//	ErrProductNotFound                = errors.New("2005: Product not found")
//	ErrFailedToValidateProductRequest = errors.New("2006: Failed to validate product request")
//
//	ErrFailedToCreateStock          = errors.New("3001: Failed to create stock")
//	ErrFailedToUpdateStock          = errors.New("3002: Failed to update stock")
//	ErrFailedToDeleteStock          = errors.New("3003: Failed to delete stock")
//	ErrFailedToRetrieveStocks       = errors.New("3004: Failed to retrieve stocks")
//	ErrStockNotFound                = errors.New("3005: Stock not found")
//	ErrFailedToValidateStockRequest = errors.New("3006: Failed to validate stock request")
//
//	ErrFailedToCreateTransaction          = errors.New("4001: Failed to create transaction")
//	ErrFailedToUpdateTransaction          = errors.New("4002: Failed to update transaction")
//	ErrFailedToDeleteTransaction          = errors.New("4003: Failed to delete transaction")
//	ErrFailedToRetrieveTransactions       = errors.New("4004: Failed to retrieve transactions")
//	ErrTransactionNotFound                = errors.New("4005: Transaction not found")
//	ErrFailedToValidateTransactionRequest = errors.New("4006: Failed to validate transaction request")
//)
