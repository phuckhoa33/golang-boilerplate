package routes

import (
	"golang-boilerplate/server"
	"golang-boilerplate/server/handlers"

	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func ConfigureRoutes(server *server.Server) {
	// Create handlers
	authHandler := handlers.NewAuthHandler(server)

	// Add middleware
	server.Echo.Use(middleware.Logger())

	v1 := server.Echo.Group("/api/v1")
	v1.GET("/login", authHandler.Login)

	// Integrate swagger to api
	server.Echo.GET("/docs/*", echoSwagger.WrapHandler)
}
