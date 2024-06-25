package product

import (
	"encoding/json"
	"github.com/kizmey/order_management_system/pkg/interface/modelReq"
	_productService "github.com/kizmey/order_management_system/pkg/service/product"
	"github.com/kizmey/order_management_system/server/httpEchoServer/custom"
	customTracer "github.com/kizmey/order_management_system/tracer"
	"github.com/labstack/echo/v4"
	"go.opentelemetry.io/otel/attribute"
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
	ctx, sp := tracer.Start(pctx.Request().Context(), "productCreateController")
	defer sp.End()

	productReq := new(modelReq.Product)

	validatingContext := custom.NewCustomEchoRequest(pctx)
	if err := validatingContext.BindAndValidate(productReq); err != nil {
		return custom.Error(pctx, http.StatusBadRequest, custom.ErrFailedToValidateProductRequest)
	}
	product, err := c.productService.Create(ctx, productReq)
	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, custom.ErrFailedToCreateProduct)
	}

	customTracer.SetSubAttributesWithJson(product, sp)

	return pctx.JSON(http.StatusCreated, product)
}

func (c *productController) FindAll(pctx echo.Context) error {
	ctx, sp := tracer.Start(pctx.Request().Context(), "productFindAllController")
	defer sp.End()

	productListingResult, err := c.productService.FindAll(ctx)
	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, custom.ErrFailedToRetrieveProducts)
	}

	productJSON, err := json.Marshal(productListingResult)
	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, custom.ErrFailedToRetrieveProducts)
	}

	sp.SetAttributes(
		attribute.String("product.listing", string(productJSON)),
	)

	return pctx.JSON(http.StatusOK, productListingResult)
}

func (c *productController) FindByID(pctx echo.Context) error {
	ctx, sp := tracer.Start(pctx.Request().Context(), "productFindByIdController")
	defer sp.End()

	id := pctx.Param("id")

	product, err := c.productService.FindByID(ctx, id)
	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, custom.ErrProductNotFound)
	}
	return pctx.JSON(http.StatusOK, product)
}

func (c *productController) Update(pctx echo.Context) error {
	ctx, sp := tracer.Start(pctx.Request().Context(), "productUpdateController")
	defer sp.End()

	id := pctx.Param("id")

	productReq := new(modelReq.Product)

	validatingContext := custom.NewCustomEchoRequest(pctx)
	if err := validatingContext.BindAndValidate(productReq); err != nil {
		return custom.Error(pctx, http.StatusBadRequest, custom.ErrFailedToValidateProductRequest)
	}
	product, err := c.productService.Update(ctx, id, productReq)
	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, custom.ErrFailedToUpdateProduct)
	}

	return pctx.JSON(http.StatusOK, product)
}

func (c *productController) Delete(pctx echo.Context) error {
	ctx, sp := tracer.Start(pctx.Request().Context(), "productDeleteController")
	defer sp.End()

	id := pctx.Param("id")

	product, err := c.productService.Delete(ctx, id)
	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, custom.ErrFailedToDeleteProduct)
	}

	return pctx.JSON(http.StatusOK, product)
}
