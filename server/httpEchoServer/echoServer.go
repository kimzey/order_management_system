package httpEchoServer

import (
	"fmt"
	"github.com/kizmey/order_management_system/config"
	logger "github.com/kizmey/order_management_system/logs"
	"github.com/kizmey/order_management_system/server"
	customMiddleware "github.com/kizmey/order_management_system/server/httpEchoServer/middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"

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

	server.InitMetrics()

	return serverEcho
}

func (s *echoServer) Start() {
	s.app.GET("/v1/health", s.healthCheck)
	s.app.GET("/metrics", echo.WrapHandler(promhttp.Handler()))

	s.app.Use(middleware.Recover())
	s.app.Use(middleware.Logger())
	s.app.Use(LoggerMiddleware)

	s.app.Use(customMiddleware.TracingMiddleware)
	//s.app.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
	//	Output: logger.LogFile,
	//}))

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

func LoggerMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		start := time.Now()

		if err := next(c); err != nil {
			c.Error(err)
		}

		fields := logrus.Fields{
			"method":   c.Request().Method,
			"path":     c.Path(),
			"query":    c.QueryString(),
			"remoteIP": c.RealIP(),
		}

		// Log HTTP request
		logger.LogInfo("HTTP request", fields)

		// Log HTTP response
		fields["status"] = c.Response().Status
		fields["latency"] = time.Since(start).Seconds()

		logger.LogInfo("HTTP response", fields)

		return nil
	}
}
