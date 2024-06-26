package custom

import (
	"errors"
	"fmt"
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
		//fmt.Errorf("invalid validate request : %s", err.Error())
		return errors.New(fmt.Sprintf("invalid request "))
	}
	if err := r.validator.Struct(obj); err != nil {
		return errors.New(fmt.Sprintf("invalid validate request "))
	}
	return nil
}
