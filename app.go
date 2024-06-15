package application

import (
	"golang-boilerplate/config"
	"golang-boilerplate/server"
	"golang-boilerplate/server/routes"
	"log"
)

func Start(config *config.Config) {
	// Initiate app server
	app := server.NewServer(config)

	// Configurate route for app
	routes.ConfigureRoutes(app)

	// Check port is existed
	err := app.Start(config.HTTP.AppPort)

	if err != nil {
		log.Fatal("Port already used")
	}
}
