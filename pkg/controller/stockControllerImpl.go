package controller

import (
	"fmt"
	"github.com/kizmey/order_management_system/pkg/interface/modelReq"
	_StockService "github.com/kizmey/order_management_system/pkg/service"
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
	stockReq := new(modelReq.Stock)

	validatingContext := custom.NewCustomEchoRequest(pctx)
	if err := validatingContext.BindAndValidate(stockReq); err != nil {
		return custom.Error(pctx, http.StatusBadRequest, err)
	}

	address, err := c.stockService.Create(stockReq)

	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, err)
	}

	return pctx.JSON(http.StatusCreated, address)

}

func (c *stockControllerImpl) FindAll(pctx echo.Context) error {
	stockListingResult, err := c.stockService.FindAll()

	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, err)
	}
	return pctx.JSON(http.StatusOK, stockListingResult)
}

func (c *stockControllerImpl) CheckStockByProductId(pctx echo.Context) error {
	id := pctx.Param("id")
	fmt.Println(id)
	//id, err := custom.CheckParamId(pctx)
	//
	//if err != nil {
	//	return custom.Error(pctx, http.StatusBadRequest, err)
	//}

	stockListingResult, err := c.stockService.CheckStockByProductId(id)

	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, err)
	}
	return pctx.JSON(http.StatusOK, stockListingResult)
}

func (c *stockControllerImpl) Update(pctx echo.Context) error {
	id := pctx.Param("id")
	fmt.Println(id)
	//stockid, err := custom.CheckParamId(pctx)
	//
	//if err != nil {
	//	return custom.Error(pctx, http.StatusBadRequest, err)
	//}

	stockReq := new(modelReq.Stock)

	validatingContext := custom.NewCustomEchoRequest(pctx)
	if err := validatingContext.BindAndValidate(stockReq); err != nil {
		return custom.Error(pctx, http.StatusBadRequest, err)
	}

	stockupdate, err := c.stockService.Update(id, stockReq)

	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, err)

	}

	return pctx.JSON(http.StatusCreated, stockupdate)
}

func (c *stockControllerImpl) Delete(pctx echo.Context) error {
	id := pctx.Param("id")
	fmt.Println(id)
	//stockid, err := custom.CheckParamId(pctx)
	//
	//if err != nil {
	//	return custom.Error(pctx, http.StatusBadRequest, err)
	//}

	stock, err := c.stockService.Delete(id)

	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, err)
	}
	return pctx.JSON(http.StatusOK, stock)
}
