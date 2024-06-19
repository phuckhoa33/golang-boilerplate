package main

import (
	application "golang-boilerplate"
	"golang-boilerplate/config"
)

func main() {
	config := config.NewConfig()

	// Start app
	application.Start(config)
}
