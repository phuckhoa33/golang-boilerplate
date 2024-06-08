package routes

import (
	"golang-boilerplate/server"
	"golang-boilerplate/server/handlers"

	echoSwagger "github.com/swaggo/echo-swagger"
)

func ConfigureRoutes(server *server.Server) {
	// Create handlers
	authHandler := handlers.NewAuthHandler(server)

	// v1 := server.Echo.Group("/api/v1")

	server.Echo.GET("/login", authHandler.Login)

	// Config swagger
	server.Echo.GET("/swagger/*", echoSwagger.WrapHandler)
}
