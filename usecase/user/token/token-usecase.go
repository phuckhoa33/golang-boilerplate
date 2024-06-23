package usecase

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

type TokenUsecaseWrapper interface {
	CreateAccessToken(user *models.User) (accessToken string, exp int64, err error)
	CreateRefreshToken(user *models.User) (t string, err error)
}

type UserTokenUsecase struct {
	config *config.Config
}

func NewTokenUsecase(config *config.Config) *UserTokenUsecase {
	return &UserTokenUsecase{
		config: config,
	}
}
