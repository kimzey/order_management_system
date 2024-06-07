package controller

import (
	"github.com/go-playground/validator/v10"
	"github.com/kizmey/order_management_system/entities"
	_orderService "github.com/kizmey/order_management_system/pkg/order/service"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type orderControllerImpl struct {
	orderService _orderService.OrderService
}

func NewOrderControllerImpl(orderService _orderService.OrderService) OrderController {
	return &orderControllerImpl{orderService}
}

func (c *orderControllerImpl) Create(pctx echo.Context) error {
	newOrderReq := new(entities.Order)

	if err := pctx.Bind(newOrderReq); err != nil {
		return pctx.JSON(http.StatusBadRequest, err.Error())
	}

	validatorInit := validator.New()
	if err := validatorInit.Struct(newOrderReq); err != nil {
		return pctx.JSON(http.StatusBadRequest, err.Error())
	}
	newOrderRes, err := c.orderService.Create(newOrderReq)
	if err != nil {
		return pctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return pctx.JSON(http.StatusCreated, newOrderRes)
}

func (c *orderControllerImpl) ChangeStatusNext(pctx echo.Context) error {
	orderId, err := strconv.ParseUint(pctx.Param("id"), 0, 64)

	if err != nil {
		return pctx.JSON(http.StatusInternalServerError, err.Error())
	}

	order, err := c.orderService.ChangeStatusNext(orderId)
	if err != nil {
		return pctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return pctx.JSON(http.StatusOK, order)
}

func (c *orderControllerImpl) FindAll(pctx echo.Context) error {
	orderListingResult, err := c.orderService.FindAll()
	if err != nil {
		return pctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return pctx.JSON(http.StatusOK, orderListingResult)
}

func (c *orderControllerImpl) FindByID(pctx echo.Context) error {
	id, err := strconv.ParseUint(pctx.Param("id"), 0, 64)

	if err != nil {
		return pctx.JSON(http.StatusInternalServerError, err.Error())
	}
	order, err := c.orderService.FindByID(id)
	if err != nil {
		return pctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return pctx.JSON(http.StatusOK, order)
}
func (c *orderControllerImpl) Update(pctx echo.Context) error {
	id, err := strconv.ParseUint(pctx.Param("id"), 0, 64)
	if err != nil {
		return pctx.JSON(http.StatusInternalServerError, err.Error())
	}
	orderReq := new(entities.Order)
	if err := pctx.Bind(orderReq); err != nil {
		return pctx.JSON(http.StatusBadRequest, err.Error())
	}

	order, err := c.orderService.Update(id, orderReq)
	if err != nil {
		return pctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return pctx.JSON(http.StatusOK, order)
}

func (c *orderControllerImpl) Delete(pctx echo.Context) error {
	id, err := strconv.ParseUint(pctx.Param("id"), 0, 64)
	if err != nil {
		return pctx.JSON(http.StatusInternalServerError, err.Error())
	}

	err = c.orderService.Delete(id)
	if err != nil {
		return pctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return pctx.JSON(http.StatusOK, "deleted")
}

func (c *orderControllerImpl) ChageStatusDone(pctx echo.Context) error {
	orderId, err := strconv.ParseUint(pctx.Param("id"), 0, 64)

	if err != nil {
		return pctx.JSON(http.StatusInternalServerError, err.Error())
	}

	order, err := c.orderService.ChageStatusDone(orderId)
	if err != nil {
		return pctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return pctx.JSON(http.StatusOK, order)
}
