package main

import (
	application "golang-boilerplate"
	"golang-boilerplate/config"
)

//	@title			Golang Boilerplate API
//	@version		1.0
//	@description	This is a boilerplate for building RESTful APIs in Golang.
//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization

func main() {
	config := config.NewConfig()

	// Start app
	application.Start(config)
}
