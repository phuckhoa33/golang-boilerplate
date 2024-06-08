package handlers

import (
	"golang-boilerplate/server"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	server *server.Server
}

func NewAuthHandler(server *server.Server) *AuthHandler {
	return &AuthHandler{server: server}
}

func (authHandler AuthHandler) Login(context echo.Context) error {
	return nil
}
