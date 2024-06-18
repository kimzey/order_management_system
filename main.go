// main.go
package main

import (
	"github.com/kizmey/order_management_system/config"
	"github.com/kizmey/order_management_system/database"
	logger "github.com/kizmey/order_management_system/logs"
	"github.com/kizmey/order_management_system/pkg"
	serverPkg "github.com/kizmey/order_management_system/server/httpEchoServer"
	"github.com/sirupsen/logrus"
)

func main() {
	// Initialize Logger
	logger.InitLogger()
	defer logger.LogFile.Close()

	// Initialize Loki client
	//logger.InitLokiClient()

	// Log to file and Loki
	fields := logrus.Fields{"module": "main", "function": "main"}
	logger.LogInfo("Service started", fields)
	//logger.LogToLoki(logrus.InfoLevel, "Service started", fields)

	conf := config.GettingConfig()
	db := database.NewPostgresDatabase(conf.Database)

	usecases := pkg.InitUsecase(db)

	server := serverPkg.NewEchoServer(conf, usecases)
	server.Start()
}
