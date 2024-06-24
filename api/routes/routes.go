package routes

import (
	user_auth_route "golang-boilerplate/api/routes/user/auth"
	"golang-boilerplate/server"
	"log"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func ConfigureRoutes(server *server.Server) {

	// Version 1 of router
	v1 := server.Gin.Group(server.Config.App.AppApiPrefix + "/v1")

	// Configure for authentication route
	user_auth_route.ConfigureUserAuthRoute(server, v1)

	// Configurate path of swagger
	server.Gin.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Talk out swagger host
	log.Printf("Swagger listening and serving at http://%s:%s/docs/index.html", server.Config.App.AppHost, server.Config.App.AppPort)

}
