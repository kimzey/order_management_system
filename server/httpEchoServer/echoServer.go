package httpEchoServer

import (
	"fmt"
	"github.com/kizmey/order_management_system/config"
	"github.com/kizmey/order_management_system/server"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"net/http"

	"github.com/kizmey/order_management_system/pkg"
)

type echoServer struct {
	app     *echo.Echo
	conf    *config.Config
	usecase *pkg.Usecase
}

func NewEchoServer(conf *config.Config, usecase *pkg.Usecase) server.Server {
	echoApp := echo.New()
	echoApp.Logger.SetLevel(log.DEBUG)

	serverEcho := &echoServer{
		app:     echoApp,
		conf:    conf,
		usecase: usecase,
	}

	return serverEcho
}

func (s *echoServer) Start() {
	s.app.GET("/v1/health", s.healthCheck)
	s.app.Use(middleware.Recover())
	s.app.Use(middleware.Logger())

	s.initStockRouter()
	s.initProductRouter()
	s.initOrderRouter()
	s.inittransactionRouter()

	s.httpListening()
}

func (s *echoServer) httpListening() {
	Url := fmt.Sprintf(":%d", s.conf.Server.Port)

	err := s.app.Start(Url)

	if err != nil {
		s.app.Logger.Fatalf("Error: %s", err.Error())
	}
}

// path : /v1/health method : GET FOR check server
func (s *echoServer) healthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "Ok")
}
