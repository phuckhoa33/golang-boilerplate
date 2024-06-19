package application

import (
	"golang-boilerplate/api/routes"
	"golang-boilerplate/config"
	"golang-boilerplate/server"
	"log"
)

func Start(config *config.Config) {
	// Initiate server
	app := server.NewServer(config)

	// Configure routes
	routes.ConfigureRoutes(app)

	// Run port
	err := app.Run(config.App.AppPort)

	// Check err when run port
	if err != nil {
		log.Fatal("Port already used, please change to another port")
	}
}
