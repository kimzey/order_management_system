package stock

import (
	"github.com/labstack/echo/v4"
	"go.opentelemetry.io/otel"
)

type StockController interface {
	Create(pctx echo.Context) error

	FindAll(pctx echo.Context) error

	CheckStockByProductId(pctx echo.Context) error

	Update(pctx echo.Context) error

	Delete(pctx echo.Context) error
}

var tracer = otel.Tracer("StockController")
