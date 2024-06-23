package controllers

import (
	requests "golang-boilerplate/domain/requests/user/auth"
	responses "golang-boilerplate/domain/responses/user/auth"
	"golang-boilerplate/models"
	respositories "golang-boilerplate/respositories/postgresql"
	"golang-boilerplate/server"
	usecase "golang-boilerplate/usecase/user/token"

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
//	@Tags			User Actions
//	@Accept			json
//	@Produce		json
//	@Param			params	body		requests.LoginRequest	true	"User's credentials"
//	@Success		200		{object}	responses.LoginResponse
//	@Failure		401		{object}	responses.Error
//	@Router			/login [post]
func (controller *UserAuthV1Controller) Login(context *gin.Context) error {
	loginRequest := new(requests.LoginRequest)

	user := models.User{}
	userRepository := respositories.NewUserRepository(controller.server.DB)

	userRepository.GetUserByEmail(&user, loginRequest.Email)

	tokenUsecase := usecase.NewTokenUsecase(controller.server.Config)
	accessToken, exp, err := tokenUsecase.CreateAccessToken(&user)
	if err != nil {
		return err
	}
	refreshToken, err := tokenUsecase.CreateRefreshToken(&user)
	if err != nil {
		return err
	}
	res := responses.NewLoginResponse(accessToken, refreshToken, exp)
	return responses.
}
