package custom

import (
	"net/http"
	"sync"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type EchoRequest interface {
	BindAndValidate(obj any) error
}

type customEchoRequest struct {
	ctx       echo.Context
	validator *validator.Validate
}

var (
	once              sync.Once
	validatorInstance *validator.Validate
)

func NewCustomEchoRequest(echoRequest echo.Context) EchoRequest {
	once.Do(func() {
		validatorInstance = validator.New()
	})

	return &customEchoRequest{
		ctx:       echoRequest,
		validator: validatorInstance,
	}
}

func (r *customEchoRequest) BindAndValidate(obj any) error {
	if err := r.ctx.Bind(obj); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload").SetInternal(err)
	}

	if err := r.validator.Struct(obj); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Validation failed").SetInternal(err)
	}

	return nil
}

//func Error(ctx echo.Context, statusCode int, err error) error {
//	return ctx.JSON(statusCode, map[string]interface{}{
//		"error":   err.Error(),
//		"message": http.StatusText(statusCode),
//	})
//}
