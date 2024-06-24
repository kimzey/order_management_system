// main.go
package main

import (
	"github.com/kizmey/order_management_system/config"
	"github.com/kizmey/order_management_system/database"
	logger "github.com/kizmey/order_management_system/logs"
	"github.com/kizmey/order_management_system/pkg"
	serverPkg "github.com/kizmey/order_management_system/server/httpEchoServer"
	"github.com/kizmey/order_management_system/tracer"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	// Initialize Logger
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
	err := customTracer.InitOpenTelemetry()
	if err != nil {
		logger.LogError("Failed to initialize OpenTelemetry"+err.Error(), fields)
	}

	server := serverPkg.NewEchoServer(conf, usecases)
	server.Start()
}
