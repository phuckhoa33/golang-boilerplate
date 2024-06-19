package application

import (
	"golang-boilerplate/config"
	"golang-boilerplate/internal/server"
	"log"
)

func Start(config *config.Config) {
	app := server.NewServer(config)

	err := app.Run(config.App.AppPort)

	if err != nil {
		log.Fatal("Port already used, please change to another port")
	}
}
