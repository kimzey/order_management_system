package controller

import "github.com/labstack/echo/v4"

type OrderController interface {
	Create(pctx echo.Context) error
	ChangeStatusNext(pctx echo.Context) error
	ChageStatusDone(pctx echo.Context) error
}
