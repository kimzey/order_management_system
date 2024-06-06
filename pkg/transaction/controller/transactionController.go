package controller

import "github.com/labstack/echo/v4"

type TransactionController interface {
	Create(ctx echo.Context) error
	FindAll(ctx echo.Context) error
}
