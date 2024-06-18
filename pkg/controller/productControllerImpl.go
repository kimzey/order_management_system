package controller

import (
	"github.com/kizmey/order_management_system/pkg/interface/modelReq"
	_productService "github.com/kizmey/order_management_system/pkg/service"
	"github.com/kizmey/order_management_system/server/httpEchoServer/custom"
	"github.com/labstack/echo/v4"
	"net/http"
)

type productController struct {
	productService _productService.ProductService
}

func NewProductControllerImpl(productService _productService.ProductService) ProductController {
	return &productController{
		productService: productService,
	}
}

func (c *productController) Create(pctx echo.Context) error {
	productReq := new(modelReq.Product)

	validatingContext := custom.NewCustomEchoRequest(pctx)
	if err := validatingContext.BindAndValidate(productReq); err != nil {
		return custom.Error(pctx, http.StatusBadRequest, err)
	}
	product, err := c.productService.Create(productReq)
	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, err)
	}
	return pctx.JSON(http.StatusCreated, product)

}

func (c *productController) FindAll(pctx echo.Context) error {
	productListingResult, err := c.productService.FindAll()
	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, err)
	}
	return pctx.JSON(http.StatusOK, productListingResult)
}

func (c *productController) FindByID(pctx echo.Context) error {
	id := pctx.Param("id")

	product, err := c.productService.FindByID(id)
	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, err)
	}
	return pctx.JSON(http.StatusOK, product)
}

func (c *productController) Update(pctx echo.Context) error {
	id := pctx.Param("id")

	productReq := new(modelReq.Product)

	validatingContext := custom.NewCustomEchoRequest(pctx)
	if err := validatingContext.BindAndValidate(productReq); err != nil {
		return custom.Error(pctx, http.StatusBadRequest, err)
	}
	product, err := c.productService.Update(id, productReq)
	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, err)
	}

	return pctx.JSON(http.StatusOK, product)
}

func (c *productController) Delete(pctx echo.Context) error {
	id := pctx.Param("id")

	product, err := c.productService.Delete(id)
	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, err)
	}

	return pctx.JSON(http.StatusOK, product)
}
