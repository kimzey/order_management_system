package controller

import (
	"github.com/go-playground/validator/v10"
	"github.com/kizmey/order_management_system/entities"
	_StockService "github.com/kizmey/order_management_system/pkg/stock/service"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type stockControllerImpl struct {
	stockService _StockService.StockService
}

func NewStockControllerImpl(stockController _StockService.StockService) StockController {
	return &stockControllerImpl{stockService: stockController}
}

func (c *stockControllerImpl) Create(pctx echo.Context) error {
	stockReq := new(entities.Stock)

	if err := pctx.Bind(stockReq); err != nil {
		return err
	}
	//fmt.Println("addressReq: ", addressReq)
	validatorInit := validator.New()
	if err := validatorInit.Struct(stockReq); err != nil {
		return err
	}
	//fmt.Println("check")
	address, err := c.stockService.Create(stockReq)

	if err != nil {
		return err
	}

	return pctx.JSON(http.StatusCreated, address)

}

func (c *stockControllerImpl) FindAll(pctx echo.Context) error {
	stockListingResult, err := c.stockService.FindAll()

	if err != nil {
		return pctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return pctx.JSON(http.StatusOK, stockListingResult)
}

func (c *stockControllerImpl) CheckStockByProductId(pctx echo.Context) error {

	productid, err := strconv.ParseUint(pctx.Param("id"), 0, 64)

	if err != nil {
		return pctx.JSON(http.StatusInternalServerError, err.Error())
	}

	stockListingResult, err := c.stockService.CheckStockByProductId(productid)

	if err != nil {
		return pctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return pctx.JSON(http.StatusOK, stockListingResult)
}

func (c *stockControllerImpl) Update(pctx echo.Context) error {
	stockid, err := strconv.ParseUint(pctx.Param("id"), 0, 64)

	if err != nil {
		return pctx.JSON(http.StatusBadRequest, err.Error())
	}

	stockReq := new(entities.Stock)
	if err := pctx.Bind(stockReq); err != nil {
		return pctx.JSON(http.StatusBadRequest, err.Error())

	}
	//fmt.Println("addressReq: ", addressReq)
	validatorInit := validator.New()
	if err := validatorInit.Struct(stockReq); err != nil {
		return pctx.JSON(http.StatusBadRequest, err.Error())

	}
	//fmt.Println("check")
	stockupdate, err := c.stockService.Update(stockid, stockReq)

	if err != nil {
		return pctx.JSON(http.StatusInternalServerError, err.Error())

	}

	return pctx.JSON(http.StatusCreated, stockupdate)
}

func (c *stockControllerImpl) Delete(pctx echo.Context) error {
	stockid, err := strconv.ParseUint(pctx.Param("id"), 0, 64)

	if err != nil {
		return pctx.JSON(http.StatusInternalServerError, err.Error())
	}

	err = c.stockService.Delete(stockid)

	if err != nil {
		return pctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return pctx.JSON(http.StatusOK, "deleted successfully")
}
