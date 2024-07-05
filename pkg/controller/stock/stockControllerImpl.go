package stock

import (
	"github.com/kizmey/order_management_system/pkg/interface/entities"
	"github.com/kizmey/order_management_system/pkg/interface/modelReq"
	"github.com/kizmey/order_management_system/pkg/interface/modelRes"
	_StockService "github.com/kizmey/order_management_system/pkg/service/stock"
	"github.com/kizmey/order_management_system/server/httpEchoServer/custom"
	"github.com/labstack/echo/v4"
	"net/http"
)

type stockControllerImpl struct {
	stockService _StockService.StockService
}

func NewStockControllerImpl(stockController _StockService.StockService) StockController {
	return &stockControllerImpl{stockService: stockController}
}

func (c *stockControllerImpl) Create(pctx echo.Context) error {
	ctx, sp := tracer.Start(pctx.Request().Context(), "stockCreateController")
	defer sp.End()

	stockReq := new(modelReq.Stock)

	validatingContext := custom.NewCustomEchoRequest(pctx)
	if err := validatingContext.BindAndValidate(stockReq); err != nil {
		return custom.Error(pctx, http.StatusBadRequest, custom.ErrFailedToValidateStockRequest)
	}

	stockEntity := c.stockReqToEntity(stockReq)
	stock, err := c.stockService.Create(ctx, stockEntity)

	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, custom.ErrFailedToCreateStock)
	}
	stockRes := c.stockEntityToRes(stock)
	return pctx.JSON(http.StatusCreated, stockRes)
}

func (c *stockControllerImpl) FindAll(pctx echo.Context) error {
	ctx, sp := tracer.Start(pctx.Request().Context(), "stockFindAllController")
	defer sp.End()

	stockListingResult, err := c.stockService.FindAll(ctx)

	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, custom.ErrFailedToRetrieveStocks)
	}

	var stockRes []modelRes.Stock
	for _, stock := range *stockListingResult {
		stockRes = append(stockRes, *c.stockEntityToRes(&stock))
	}

	return pctx.JSON(http.StatusOK, stockRes)
}

func (c *stockControllerImpl) CheckStockByProductId(pctx echo.Context) error {
	ctx, sp := tracer.Start(pctx.Request().Context(), "CheckStockByProductIdController")
	defer sp.End()

	id := pctx.Param("id")

	stock, err := c.stockService.CheckStockByProductId(ctx, id)

	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, custom.ErrStockNotFound)
	}

	stockRes := c.stockEntityToRes(stock)
	return pctx.JSON(http.StatusOK, stockRes)
}

func (c *stockControllerImpl) Update(pctx echo.Context) error {
	ctx, sp := tracer.Start(pctx.Request().Context(), "stockUpdateController")
	defer sp.End()

	id := pctx.Param("id")

	stockReq := new(modelReq.Stock)

	validatingContext := custom.NewCustomEchoRequest(pctx)
	if err := validatingContext.BindAndValidate(stockReq); err != nil {
		return custom.Error(pctx, http.StatusBadRequest, custom.ErrFailedToValidateStockRequest)
	}

	stockEntity := c.stockReqToEntity(stockReq)
	stock, err := c.stockService.Update(ctx, id, stockEntity)

	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, custom.ErrFailedToUpdateStock)
	}

	stockRes := c.stockEntityToRes(stock)
	return pctx.JSON(http.StatusCreated, stockRes)
}

func (c *stockControllerImpl) Delete(pctx echo.Context) error {
	ctx, sp := tracer.Start(pctx.Request().Context(), "stockDeleteController")
	defer sp.End()
	id := pctx.Param("id")

	stock, err := c.stockService.Delete(ctx, id)

	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, custom.ErrFailedToDeleteStock)
	}

	stockRes := c.stockEntityToRes(stock)
	return pctx.JSON(http.StatusOK, stockRes)
}

func (c *stockControllerImpl) stockReqToEntity(stockReq *modelReq.Stock) *entities.Stock {
	return &entities.Stock{
		ProductID: stockReq.ProductID,
		Quantity:  stockReq.Quantity,
	}
}

func (c *stockControllerImpl) stockEntityToRes(stock *entities.Stock) *modelRes.Stock {
	return &modelRes.Stock{
		StockID:   stock.StockID,
		ProductID: stock.ProductID,
		Quantity:  stock.Quantity,
	}
}
