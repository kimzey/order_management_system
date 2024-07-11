package product

import (
	"errors"
	"github.com/kizmey/order_management_system/pkg/interface/entities"
	"github.com/kizmey/order_management_system/pkg/interface/modelReq"
	"github.com/kizmey/order_management_system/pkg/interface/modelRes"
	_productService "github.com/kizmey/order_management_system/pkg/service/product"
	"github.com/kizmey/order_management_system/server/httpEchoServer/custom"
	"github.com/labstack/echo/v4"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"net/http"
	"reflect"
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
	product := c.productReqToEntity(productReq)
	product, err := c.productService.Create(ctx, product)
	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, custom.ErrFailedToCreateProduct)
	}

	productRes := c.productEntityToRes(product)
	c.SetSubAttributesWithJson(productRes, sp)
	return pctx.JSON(http.StatusCreated, productRes)
}

func (c *productController) FindAll(pctx echo.Context) error {
	ctx, sp := tracer.Start(pctx.Request().Context(), "productFindAllController")
	defer sp.End()

	productListingResult, err := c.productService.FindAll(ctx)
	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, custom.ErrFailedToRetrieveProducts)
	}

	productsRes := make([]modelRes.Product, 0)
	for _, product := range *productListingResult {
		productsRes = append(productsRes, *c.productEntityToRes(&product))
	}

	c.SetSubAttributesWithJson(productsRes, sp)
	return pctx.JSON(http.StatusOK, productsRes)
}

func (c *productController) FindByID(pctx echo.Context) error {
	ctx, sp := tracer.Start(pctx.Request().Context(), "productFindByIdController")
	defer sp.End()

	id := pctx.Param("id")

	product, err := c.productService.FindByID(ctx, id)
	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, custom.ErrProductNotFound)
	}

	productRes := c.productEntityToRes(product)
	c.SetSubAttributesWithJson(productRes, sp)

	return pctx.JSON(http.StatusOK, productRes)
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

	product := c.productReqToEntity(productReq)
	product, err := c.productService.Update(ctx, id, product)
	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, custom.ErrFailedToUpdateProduct)
	}

	productRes := c.productEntityToRes(product)

	c.SetSubAttributesWithJson(productRes, sp)
	return pctx.JSON(http.StatusOK, productRes)
}

func (c *productController) Delete(pctx echo.Context) error {
	ctx, sp := tracer.Start(pctx.Request().Context(), "productDeleteController")
	defer sp.End()

	id := pctx.Param("id")

	product, err := c.productService.Delete(ctx, id)
	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, custom.ErrFailedToDeleteProduct)
	}

	productRes := c.productEntityToRes(product)
	c.SetSubAttributesWithJson(productRes, sp)

	return pctx.JSON(http.StatusOK, productRes)
}

func (c *productController) productReqToEntity(product *modelReq.Product) *entities.Product {
	return &entities.Product{
		ProductName:  product.ProductName,
		ProductPrice: product.ProductPrice,
	}

}

func (c *productController) productEntityToRes(product *entities.Product) *modelRes.Product {
	return &modelRes.Product{
		ProductID:    product.ProductID,
		ProductName:  product.ProductName,
		ProductPrice: product.ProductPrice,
	}

}

func (c *productController) SetSubAttributesWithJson(obj any, sp trace.Span) {
	if products, ok := obj.([]modelRes.Product); ok {
		var productIDs []string
		var productNames []string
		var productPrices []int
		var quantities []int

		for _, product := range products {
			productIDs = append(productIDs, product.ProductID)
			productNames = append(productNames, product.ProductName)
			productPrices = append(productPrices, int(product.ProductPrice))
			quantities = append(quantities, int(product.Quantity))
		}

		sp.SetAttributes(
			attribute.StringSlice("ProductID", productIDs),
			attribute.StringSlice("ProductName", productNames),
			attribute.IntSlice("ProductPrice", productPrices),
			attribute.IntSlice("Quantity", quantities),
		)
	} else if product, ok := obj.(*modelRes.Product); ok {
		sp.SetAttributes(
			attribute.String("ProductID", product.ProductID),
			attribute.String("ProductName", product.ProductName),
			attribute.Int("ProductPrice", int(product.ProductPrice)),
			attribute.Int("Quantity", int(product.Quantity)),
		)
	} else {
		sp.RecordError(errors.New("invalid type" + reflect.TypeOf(obj).String()))
	}
}
