// main.go
package main

import (
	"github.com/kizmey/order_management_system/config"
	"github.com/kizmey/order_management_system/database"
	logger "github.com/kizmey/order_management_system/observability/logs"
	customTracer "github.com/kizmey/order_management_system/observability/tracer"
	"github.com/kizmey/order_management_system/pkg"
	serverPkg "github.com/kizmey/order_management_system/server/httpEchoServer"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	//Initialize Logger
	logger.InitLogger()
	defer func(LogFile *os.File) {
		err := LogFile.Close()
		if err != nil {
		}
	}(logger.LogFile)

	fields := logrus.Fields{"module": "main", "function": "main"}
	logger.LogInfo("Service started", fields)

	conf := config.GettingConfig()
	db := database.NewPostgresDatabase(conf.Database)

	usecases := pkg.InitUsecase(db)
	err := customTracer.InitOpenTelemetry(conf.Observability)
	if err != nil {
		logger.LogError("Failed to initialize OpenTelemetry"+err.Error(), fields)
	}

	server := serverPkg.NewEchoServer(conf, usecases)
	server.Start()
}
