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

		fields := logrus.Fields{
			"method":   c.Request().Method,
			"path":     c.Path(),
			"query":    c.QueryString(),
			"remoteIP": c.RealIP(),
		}

		// Log HTTP request
		logger.LogInfo("HTTP request", fields)

		if err := next(c); err != nil {
			c.Error(err)
		}

		// Log HTTP response
		fields["status"] = c.Response().Status
		fields["latency"] = time.Since(start).Seconds()

		logger.LogInfo("HTTP response", fields)

		return nil
	}
}

//package customMiddleware
//
//import (
//"bytes"
//"io/ioutil"
//"net/http"
//"time"
//
//logger "github.com/kizmey/order_management_system/logs"
//"github.com/labstack/echo/v4"
//"github.com/sirupsen/logrus"
//)
//
//func LoggerMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
//	return func(c echo.Context) error {
//		start := time.Now()
//
//		// Read and log request details
//		reqBody, err := ioutil.ReadAll(c.Request().Body)
//		if err != nil {
//			logger.LogError("Failed to read request body", logrus.Fields{"error": err.Error()})
//			return err
//		}
//		c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(reqBody)) // Restore body
//
//		fields := logrus.Fields{
//			"method":      c.Request().Method,
//			"path":        c.Path(),
//			"query":       c.QueryString(),
//			"remoteIP":    c.RealIP(),
//			"user-agent":  c.Request().UserAgent(),
//			"headers":     c.Request().Header,
//			"requestBody": string(reqBody),
//		}
//
//		logger.LogInfo("HTTP request", fields)
//
//		// Capture response details
//		resBody := &bytes.Buffer{}
//		writer := &responseLogger{ResponseWriter: c.Response().Writer, body: resBody}
//		c.Response().Writer = writer
//
//		if err := next(c); err != nil {
//			c.Error(err)
//		}
//
//		// Log response details
//		fields["status"] = c.Response().Status
//		fields["latency"] = time.Since(start).Seconds()
//		fields["responseBody"] = resBody.String()
//		fields["responseSize"] = c.Response().Size
//
//		logger.LogInfo("HTTP response", fields)
//
//		return nil
//	}
//}
//
//type responseLogger struct {
//	http.ResponseWriter
//	body *bytes.Buffer
//}
//
//func (rl *responseLogger) Write(b []byte) (int, error) {
//	rl.body.Write(b)
//	return rl.ResponseWriter.Write(b)
//}
//
//func (rl *responseLogger) WriteHeader(statusCode int) {
//	rl.ResponseWriter.WriteHeader(statusCode)
//}
