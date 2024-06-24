package user_auth_route

import (
	user_auth_v1_controller "golang-boilerplate/api/controllers/v1/user"
	"golang-boilerplate/server"

	"github.com/gin-gonic/gin"
)

func ConfigureUserAuthRoute(server *server.Server, route *gin.RouterGroup) {
	// Import controller
	authController := user_auth_v1_controller.NewUserAuthV1Controller(server)

	// Config route
	userAuthRoute := route.Group("/user")
	{
		userAuthRoute.POST("/login", authController.Login)
	}
}
