package controller

import (
	"errors"
	"fmt"
	"github.com/kizmey/order_management_system/pkg/interface/modelReq"
	_orderService "github.com/kizmey/order_management_system/pkg/service"
	"github.com/kizmey/order_management_system/server/httpEchoServer/custom"
	"github.com/labstack/echo/v4"
	"net/http"
)

type orderControllerImpl struct {
	orderService _orderService.OrderService
}

func NewOrderControllerImpl(orderService _orderService.OrderService) OrderController {
	return &orderControllerImpl{orderService}
}

func (c *orderControllerImpl) Create(pctx echo.Context) error {
	newOrderReq := new(modelReq.Order)

	validatingContext := custom.NewCustomEchoRequest(pctx)
	if err := validatingContext.BindAndValidate(newOrderReq); err != nil {
		return custom.Error(pctx, http.StatusBadRequest, err)
	}

	newOrderRes, err := c.orderService.Create(newOrderReq)
	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, err)
	}
	return pctx.JSON(http.StatusCreated, newOrderRes)
}

func (c *orderControllerImpl) ChangeStatusNext(pctx echo.Context) error {

	orderId := pctx.Param("id")
	if orderId == "" {
		return custom.Error(pctx, http.StatusBadRequest, errors.New("order id not found"))
	}

	//orderId, err := custom.CheckParamId(pctx)
	//if err != nil {
	//	return custom.Error(pctx, http.StatusBadRequest, err)
	//}

	order, err := c.orderService.ChangeStatusNext(orderId)
	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, err)
	}
	return pctx.JSON(http.StatusOK, order)
}

func (c *orderControllerImpl) FindAll(pctx echo.Context) error {
	orderListingResult, err := c.orderService.FindAll()

	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, err)
	}
	return pctx.JSON(http.StatusOK, orderListingResult)
}

func (c *orderControllerImpl) FindByID(pctx echo.Context) error {
	id := pctx.Param("id")
	fmt.Println(id)

	//id, err := custom.CheckParamId(pctx)
	//if err != nil {
	//	return custom.Error(pctx, http.StatusBadRequest, err)
	//}

	order, err := c.orderService.FindByID(id)
	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, err)
	}
	return pctx.JSON(http.StatusOK, order)
}
func (c *orderControllerImpl) Update(pctx echo.Context) error {
	id := pctx.Param("id")
	fmt.Println(id)
	//id, err := custom.CheckParamId(pctx)
	//if err != nil {
	//	return custom.Error(pctx, http.StatusBadRequest, err)
	//}

	orderReq := new(modelReq.Order)
	validatingContext := custom.NewCustomEchoRequest(pctx)
	if err := validatingContext.BindAndValidate(orderReq); err != nil {
		return custom.Error(pctx, http.StatusBadRequest, err)
	}

	order, err := c.orderService.Update(id, orderReq)
	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, err)
	}

	return pctx.JSON(http.StatusOK, order)
}

func (c *orderControllerImpl) Delete(pctx echo.Context) error {
	id := pctx.Param("id")
	fmt.Println(id)
	//id, err := custom.CheckParamId(pctx)
	//if err != nil {
	//	return custom.Error(pctx, http.StatusBadRequest, err)
	//}

	err := c.orderService.Delete(id)
	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, err)
	}
	return pctx.JSON(http.StatusOK, "deleted")
}

func (c *orderControllerImpl) ChageStatusDone(pctx echo.Context) error {
	id := pctx.Param("id")
	fmt.Println(id)
	//id, err := custom.CheckParamId(pctx)
	//if err != nil {
	//	return custom.Error(pctx, http.StatusBadRequest, err)
	//}

	fmt.Println("test")
	order, err := c.orderService.ChageStatusDone(id)
	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, err)
	}

	return pctx.JSON(http.StatusOK, order)
}
