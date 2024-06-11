package pkg

import (
	"github.com/kizmey/order_management_system/database"
	_stockRepository "github.com/kizmey/order_management_system/pkg/repository"
	"github.com/kizmey/order_management_system/pkg/service"
)

type Usecase struct {
	TransactionService service.TransactionService
	StockService       service.StockService
	ProductService     service.ProductService
	OrderService       service.OrderService
}

func NewUsecase(
	transactionService service.TransactionService,
	stockService service.StockService,
	productService service.ProductService,
	orderService service.OrderService,
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
	orderRepo := _stockRepository.NewOrderRepositoryImpl(db)
	productRepo := _stockRepository.NewProductRepositoryImpl(db)
	stockRepo := _stockRepository.NewStockRepositoryImpl(db)
	transactionRepo := _stockRepository.NewTransactionController(db)

	// Init Service
	productService := service.NewProductServiceImpl(productRepo)
	stockService := service.NewStockServiceImpl(stockRepo)
	transactionService := service.NewTransactionServiceImpl(transactionRepo, productRepo)
	orderService := service.NewOrderServiceImpl(orderRepo, transactionRepo, stockRepo)

	usecase := NewUsecase(transactionService, stockService, productService, orderService)

	return usecase
}
