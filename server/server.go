package server

import (
	"golang-boilerplate/config"
	db "golang-boilerplate/db/postgres"
	"golang-boilerplate/services"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Server struct {
	Gin                    *gin.Engine
	DB                     *gorm.DB
	Config                 *config.Config
	ServicesInitialization *services.ServicesInitialization
}

func NewServer(config *config.Config) *Server {
	return &Server{
		Gin:                    gin.Default(),
		DB:                     db.Init(config),
		Config:                 config,
		ServicesInitialization: services.NewServicesInitialization(config),
	}
}

func (server *Server) Run(appPort string) error {
	// Configure swagger info
	InitSwaggerInfo(server.Config)

	// Run application
	return server.Gin.Run(":" + appPort)
}
