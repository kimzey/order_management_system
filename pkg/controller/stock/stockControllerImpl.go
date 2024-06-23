package stock

import (
	"fmt"
	"github.com/kizmey/order_management_system/pkg/interface/modelReq"
	_StockService "github.com/kizmey/order_management_system/pkg/service/stock"
	"github.com/kizmey/order_management_system/server/httpEchoServer/custom"
	"github.com/labstack/echo/v4"
	"go.opentelemetry.io/otel/trace"
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
		return custom.Error(pctx, http.StatusBadRequest, err)
	}

	address, err := c.stockService.Create(ctx, stockReq)

	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, err)
	}

	return pctx.JSON(http.StatusCreated, address)

}

func (c *stockControllerImpl) FindAll(pctx echo.Context) error {
	ctx, sp := tracer.Start(pctx.Request().Context(), "stockFindAllController")
	defer sp.End()

	// ตรวจสอบ Trace ID
	traceID := trace.SpanContextFromContext(ctx).TraceID()
	fmt.Println("Trace ID Controller: ", traceID)

	stockListingResult, err := c.stockService.FindAll(ctx)

	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, err)
	}
	return pctx.JSON(http.StatusOK, stockListingResult)
}

func (c *stockControllerImpl) CheckStockByProductId(pctx echo.Context) error {
	ctx, sp := tracer.Start(pctx.Request().Context(), "CheckStockByProductIdController")
	defer sp.End()

	id := pctx.Param("id")

	stockListingResult, err := c.stockService.CheckStockByProductId(ctx, id)

	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, err)
	}
	return pctx.JSON(http.StatusOK, stockListingResult)
}

func (c *stockControllerImpl) Update(pctx echo.Context) error {
	ctx, sp := tracer.Start(pctx.Request().Context(), "stockUpdateController")
	defer sp.End()

	id := pctx.Param("id")

	stockReq := new(modelReq.Stock)

	validatingContext := custom.NewCustomEchoRequest(pctx)
	if err := validatingContext.BindAndValidate(stockReq); err != nil {
		return custom.Error(pctx, http.StatusBadRequest, err)
	}

	stockupdate, err := c.stockService.Update(ctx, id, stockReq)

	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, err)

	}

	return pctx.JSON(http.StatusCreated, stockupdate)
}

func (c *stockControllerImpl) Delete(pctx echo.Context) error {
	ctx, sp := tracer.Start(pctx.Request().Context(), "stockDeleteController")
	defer sp.End()
	id := pctx.Param("id")

	stock, err := c.stockService.Delete(ctx, id)

	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, err)
	}
	return pctx.JSON(http.StatusOK, stock)
}