package product

import (
	"github.com/labstack/echo/v4"
	"go.opentelemetry.io/otel"
)

type ProductController interface {
	Create(pctx echo.Context) error
	FindAll(pctx echo.Context) error
	FindByID(pctx echo.Context) error
	Update(pctx echo.Context) error
	Delete(pctx echo.Context) error
}

var tracer = otel.Tracer("ProductController")
