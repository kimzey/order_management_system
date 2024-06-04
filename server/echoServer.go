package server

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"kizmey/intern/task/database"
	"net/http"
)

type echoServer struct {
	app *echo.Echo
	db  database.Database
}

func NewEchoServer(DB database.Database) Server {
	echoApp := echo.New()
	//echoApp.logger = setLevel(log.debug)

	server := &echoServer{
		app: echoApp,
		db:  DB}
	return server
}

func (s *echoServer) Start() {
	s.app.GET("/v1/health", s.healthCheck)

	s.httpListening()
}

func (s *echoServer) httpListening() {
	Url := fmt.Sprintf(":%d", 8080)

	err := s.app.Start(Url)

	if err != nil {
		s.app.Logger.Fatalf("Error: %s", err.Error())
	}
}

func (s *echoServer) healthCheck(c echo.Context) error {
	s.db.Connect()
	return c.String(http.StatusOK, "Ok")
}
