package user_route

import (
	user_v1_controller "golang-boilerplate/api/controllers/v1/user"
	"golang-boilerplate/server"

	"github.com/gin-gonic/gin"
)

func ConfigureUserAuthRoute(server *server.Server, route *gin.RouterGroup) {
	// Import controller
	authController := user_v1_controller.NewUserAuthV1Controller(server)

	// Config route
	userAuthRoute := route.Group("/user")
	{
		userAuthRoute.POST("/login", authController.Login)
		userAuthRoute.POST("/register", authController.Register)
		userAuthRoute.POST("/refresh", authController.RefeshToken)
		userAuthRoute.POST("/forgot-password", authController.ForgotPassword)
		userAuthRoute.GET("/check-valid-forgot-link/:token", authController.CheckValidForgotPasswordLink)
		userAuthRoute.POST("/forgot-password/:token", authController.ForgotPasswordToResetPassword)
	}
}
