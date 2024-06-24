package user_auth_v1_controller

import (
	requests "golang-boilerplate/domain/requests/user/auth"
	"golang-boilerplate/domain/responses"
	responses_user_auth "golang-boilerplate/domain/responses/user/auth"
	"golang-boilerplate/models"
	respositories "golang-boilerplate/respositories/postgresql"
	"golang-boilerplate/server"
	services "golang-boilerplate/services/user/token"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserAuthV1Controller struct {
	server *server.Server
}

func NewUserAuthV1Controller(server *server.Server) *UserAuthV1Controller {
	return &UserAuthV1Controller{server: server}
}

// Login godoc
//
//	@Summary		Authenticate a user
//	@Description	Perform user login
//	@ID				user-login
//	@Tags			user.auth
//	@Accept			json
//	@Produce		json
//	@Param			params	body		requests.LoginRequest	true	"User's credentials"
//	@Success		200		{object}	responses_user_auth.LoginResponse
//	@Failure		401		{object}	responses.Error
//	@Router			/user/login [post]
func (controller *UserAuthV1Controller) Login(context *gin.Context) {
	loginRequest := new(requests.LoginRequest)

	user := models.User{}
	userRepository := respositories.NewUserRepository(controller.server.DB)

	userRepository.GetUserByEmail(&user, loginRequest.Email)

	tokenService := services.NewTokenService(controller.server.Config)
	accessToken, exp, err := tokenService.CreateAccessToken(&user)
	if err != nil {
		return
	}
	refreshToken, err := tokenService.CreateRefreshToken(&user)
	if err != nil {
		return
	}
	res := responses_user_auth.NewLoginResponse(accessToken, refreshToken, exp)
	responses.Response(context, http.StatusOK, res)
}
