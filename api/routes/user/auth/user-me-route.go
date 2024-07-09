package user_route

import (
	user_v1_controller "golang-boilerplate/api/controllers/v1/user"
	middleware "golang-boilerplate/api/middlewares"
	"golang-boilerplate/server"

	"github.com/gin-gonic/gin"
)

func ConfigureUserMeRoute(server *server.Server, route *gin.RouterGroup) {
	// Import controller
	profileController := user_v1_controller.NewUserMeV1Controller(server)

	// Configure route
	userMeRoute := route.Group("/user")
	userMeRoute.Use(middleware.AuthenticationMiddleware(server.Config))
	{
		userMeRoute.GET("/me", profileController.GetUserProfile)
		userMeRoute.PUT("/", profileController.UpdateUserInfo)
		userMeRoute.GET("/pre-signed-put-url/:objectName", profileController.GetPreSignedPutURL)
		userMeRoute.PATCH("/change-password", profileController.ChangePassword)
	}
}
