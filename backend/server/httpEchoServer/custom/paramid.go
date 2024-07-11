package custom

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"strconv"
)

func CheckParamId(echoRequest echo.Context) (uint64, error) {
	paramid := echoRequest.Param("id")
	//check error
	if paramid == "done/" {
		return 0, errors.New("param id not found")

	}

	id, err := strconv.ParseUint(paramid, 0, 64)
	if err != nil {
		return 0, errors.New(fmt.Sprintf("param is not number : %s", err.Error()))
	}

	return id, nil
}
