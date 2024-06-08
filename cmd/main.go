package main

import (
	application "golang-boilerplate"
	"golang-boilerplate/config"
	_ "golang-boilerplate/docs"
)

func main() {
	config := config.NewConfig()

	application.Start(config)
}
