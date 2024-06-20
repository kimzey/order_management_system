package pkg

import (
	"github.com/kizmey/order_management_system/database"
	"github.com/kizmey/order_management_system/pkg/repository/order"
	"github.com/kizmey/order_management_system/pkg/repository/product"
	"github.com/kizmey/order_management_system/pkg/repository/stock"
	"github.com/kizmey/order_management_system/pkg/repository/transaction"
	order2 "github.com/kizmey/order_management_system/pkg/service/order"
	product2 "github.com/kizmey/order_management_system/pkg/service/product"
	stock2 "github.com/kizmey/order_management_system/pkg/service/stock"
	transaction2 "github.com/kizmey/order_management_system/pkg/service/transaction"
)

type Usecase struct {
	TransactionService transaction2.TransactionService
	StockService       stock2.StockService
	ProductService     product2.ProductService
	OrderService       order2.OrderService
}

func NewUsecase(
	transactionService transaction2.TransactionService,
	stockService stock2.StockService,
	productService product2.ProductService,
	orderService order2.OrderService,
) *Usecase {
	return &Usecase{
		TransactionService: transactionService,
		StockService:       stockService,
		ProductService:     productService,
		OrderService:       orderService,
	}
}

func InitUsecase(db database.Database) *Usecase {

	// Init Repository

	orderRepo := order.NewOrderRepositoryImpl(db)
	productRepo := product.NewProductRepositoryImpl(db)
	stockRepo := stock.NewStockRepositoryImpl(db)
	transactionRepo := transaction.NewTransactionRepositoryImpl(db)

	// Init Service
	productService := product2.NewProductServiceImpl(productRepo)
	stockService := stock2.NewStockServiceImpl(stockRepo)
	transactionService := transaction2.NewTransactionServiceImpl(transactionRepo, productRepo)
	orderService := order2.NewOrderServiceImpl(orderRepo, transactionRepo, stockRepo)

	usecases := NewUsecase(transactionService, stockService, productService, orderService)

	return usecases
}
