package server

import (
	"golang-boilerplate/cmd/docs"
	"golang-boilerplate/config"
)

func InitSwaggerInfo(config *config.Config) {
	// Configure swagger information
	docs.SwaggerInfo.BasePath = config.App.AppApiPrefix + "/v1"
	docs.SwaggerInfo.Version = config.App.AppVersion
	docs.SwaggerInfo.Description = config.App.AppDescription
	docs.SwaggerInfo.Title = config.App.AppName
}
