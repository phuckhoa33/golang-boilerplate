package server

import (
	"golang-boilerplate/config"
	"golang-boilerplate/db"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Server struct {
	Echo   *echo.Echo
	DB     *gorm.DB
	Config *config.Config
}

func NewServer(config *config.Config) *Server {
	return &Server{
		Echo:   echo.New(),
		DB:     db.Init(config),
		Config: config,
	}
}

func (server *Server) Start(port string) error {
	return server.Echo.Start(":" + port)
}
