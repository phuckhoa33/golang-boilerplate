package application

import (
	"golang-boilerplate/config"
	"golang-boilerplate/server"
	"golang-boilerplate/server/routes"
	"log"
)

func Start(config *config.Config) {
	app := server.NewServer(config)

	routes.ConfigureRoutes(app)

	err := app.Start(config.HTTP.AppPort)

	if err != nil {
		log.Fatal("Port already used")
	}
}
