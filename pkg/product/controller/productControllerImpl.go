package controller

import (
	"github.com/go-playground/validator/v10"
	"github.com/kizmey/order_management_system/entities"
	_productService "github.com/kizmey/order_management_system/pkg/product/service"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type productController struct {
	productService _productService.ProductService
}

func NewProductController(productService _productService.ProductService) ProductController {
	return &productController{
		productService: productService,
	}
}

func (c *productController) Create(pctx echo.Context) error {
	productReq := new(entities.Product)
	if err := pctx.Bind(productReq); err != nil {
		return pctx.JSON(http.StatusBadRequest, err.Error())
	}
	//fmt.Println("productReq: ", productReq)
	validatorInit := validator.New()
	if err := validatorInit.Struct(productReq); err != nil {
		return pctx.JSON(http.StatusBadRequest, err.Error())
	}

	product, err := c.productService.Create(productReq)
	if err != nil {
		return pctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return pctx.JSON(http.StatusCreated, product)

}

func (c *productController) FindAll(pctx echo.Context) error {
	productListingResult, err := c.productService.FindAll()
	if err != nil {
		return pctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return pctx.JSON(http.StatusOK, productListingResult)
}

func (c *productController) FindByID(pctx echo.Context) error {
	productid, err := strconv.ParseUint(pctx.Param("id"), 0, 64)

	if err != nil {
		return pctx.JSON(http.StatusInternalServerError, err.Error())
	}

	product, err := c.productService.FindByID(productid)
	if err != nil {
		return pctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return pctx.JSON(http.StatusOK, product)
}

func (c *productController) Update(pctx echo.Context) error {
	productid, err := strconv.ParseUint(pctx.Param("id"), 0, 64)

	if err != nil {
		return pctx.JSON(http.StatusBadRequest, err.Error())
	}

	productReq := new(entities.Product)
	if err := pctx.Bind(productReq); err != nil {
		return pctx.JSON(http.StatusBadRequest, err.Error())
	}
	//fmt.Println("productReq: ", productReq)
	validatorInit := validator.New()
	if err := validatorInit.Struct(productReq); err != nil {
		return pctx.JSON(http.StatusBadRequest, err.Error())
	}

	product, err := c.productService.Update(productid, productReq)
	if err != nil {
		return pctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return pctx.JSON(http.StatusOK, product)
}

func (c *productController) Delete(pctx echo.Context) error {
	productid, err := strconv.ParseUint(pctx.Param("id"), 0, 64)

	if err != nil {
		return pctx.JSON(http.StatusInternalServerError, err.Error())
	}

	err = c.productService.Delete(productid)
	if err != nil {
		return pctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return pctx.JSON(http.StatusOK, "deleted")
}
