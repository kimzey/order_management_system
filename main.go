package main

import (
	"kizmey/intern/task/database"
	server_pkg "kizmey/intern/task/server"
)

func main() {
	db := database.NewPostgresDatabase()

	server := server_pkg.NewEchoServer(db)
	server.Start()

}
