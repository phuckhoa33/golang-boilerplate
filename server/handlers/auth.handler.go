package handlers

import (
	"fmt"
	"golang-boilerplate/requests"
	"golang-boilerplate/server"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	server *server.Server
}

func NewAuthHandler(server *server.Server) *AuthHandler {
	return &AuthHandler{server: server}
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
func (authHandler AuthHandler) Login(context echo.Context) error {
	loginRequest := new(requests.LoginRequest)

	fmt.Println(loginRequest)

	return nil
}
