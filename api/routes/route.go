package routes

import (
	"golang-boilerplate/server"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Helloworld(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld")
}

func ConfigureRoutes(server *server.Server) {

	v1 := server.Gin.Group(server.Config.App.AppApiPrefix + "/v1")

	{
		eg := v1.Group("/example")
		{
			eg.GET("/helloword", Helloworld)
		}
	}

	// Configurate path of swagger
	server.Gin.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Talk out swagger host
	log.Printf("Swagger listening and serving at http://%s:%s/docs/index.html", server.Config.App.AppHost, server.Config.App.AppPort)

}
