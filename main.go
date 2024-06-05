package main

import (
	"github.com/kizmey/order_management_system/config"
	"github.com/kizmey/order_management_system/database"
	serverPkg "github.com/kizmey/order_management_system/server"
)

func main() {
	conf := config.GettingConfig()
	db := database.NewPostgresDatabase(conf.Database)

	server := serverPkg.NewEchoServer(conf, db)
	server.Start()
}
