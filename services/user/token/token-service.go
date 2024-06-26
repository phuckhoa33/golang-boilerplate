package token_service

import (
	"golang-boilerplate/config"
	"golang-boilerplate/models"

	"github.com/golang-jwt/jwt/v5"
)

type JwtCustomClaims struct {
	Name string `json:"name"`
	ID   uint   `json:"id"`
	jwt.RegisteredClaims
}

type JwtCustomRefreshClaims struct {
	ID uint `json:"id"`
	jwt.RegisteredClaims
}

type TokenServiceWrapper interface {
	CreateAccessToken(user *models.User) (accessToken string, exp int64, err error)
	CreateRefreshToken(user *models.User) (t string, err error)
}

type UserTokenService struct {
	config *config.Config
}

func NewTokenService(config *config.Config) *UserTokenService {
	return &UserTokenService{
		config: config,
	}
}
