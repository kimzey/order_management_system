package customMiddleware

import (
	logger "github.com/kizmey/order_management_system/logs"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"time"
)

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
