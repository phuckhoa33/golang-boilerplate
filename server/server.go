package server

import (
	"golang-boilerplate/config"
	db "golang-boilerplate/db/postgres"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Server struct {
	Gin    *gin.Engine
	DB     *gorm.DB
	Config *config.Config
}

func NewServer(config *config.Config) *Server {
	return &Server{
		Gin:    gin.Default(),
		DB:     db.Init(config),
		Config: config,
	}
}

func (server *Server) Run(appPort string) error {
	// Configure swagger info
	InitSwaggerInfo(server.Config)

	// Run application
	return server.Gin.Run(":" + appPort)
}
