package httpEchoServer

import (
	"context"
	"fmt"
	"github.com/kizmey/order_management_system/config"
	"github.com/kizmey/order_management_system/database/migration"
	"github.com/kizmey/order_management_system/observability"
	"github.com/kizmey/order_management_system/pkg"
	"github.com/kizmey/order_management_system/server"
	customMiddleware "github.com/kizmey/order_management_system/server/httpEchoServer/middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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

	observability.InitMetrics()

	return serverEcho
}

func (s *echoServer) Start() {
	s.app.GET("/v1/health", s.healthCheck)
	s.app.GET("/metricsx", echo.WrapHandler(promhttp.Handler()))
	s.app.GET("/v1/migration", s.migration)

	s.app.Use(middleware.Recover())
	s.app.Use(middleware.Logger())

	s.app.Use(customMiddleware.LoggerMiddleware)
	s.app.Use(customMiddleware.TracingMiddleware)

	//s.app.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
	//	Output: logger.LogFile,
	//}))

	s.initStockRouter()
	s.initProductRouter()
	s.inittransactionRouter()
	s.initOrderRouter()

	// Graceful shutdown
	quitCh := make(chan os.Signal, 1)
	signal.Notify(quitCh, syscall.SIGINT, syscall.SIGTERM)
	go s.gracefullyShutdown(quitCh)

	s.httpListening()
}

func (s *echoServer) gracefullyShutdown(quitCh <-chan os.Signal) {
	ctx := context.Background()

	<-quitCh
	s.app.Logger.Infof("Shutting down service...")

	if err := s.app.Shutdown(ctx); err != nil {
		s.app.Logger.Fatalf("Error: %s", err.Error())
	}
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
	return c.String(http.StatusOK, "OK")
}

func (s *echoServer) migration(c echo.Context) error {
	migration.GettingMigration()
	return c.String(http.StatusOK, "Migration OK")
}
