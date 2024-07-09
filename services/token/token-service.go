package token_service

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang-boilerplate/config"
	"golang-boilerplate/domain/models/postgresql"
)

type JwtCustomClaims struct {
	UserID   uuid.UUID `json:"userId"`
	Username string    `json:"username"`
	jwt.RegisteredClaims
}

type JwtCustomRefreshClaims struct {
	UserID uuid.UUID `json:"userId"`
	jwt.RegisteredClaims
}

type TokenServiceWrapper interface {
	CreateAccessToken(user *postgresql.User) (accessToken string, exp int64, err error)
	CreateRefreshToken(user *postgresql.User) (t string, err error)
	VerifyToken(tokenString string) (jwt.MapClaims, error)
	CreateForgotPasswordToken(user *postgresql.User) (t string, err error)
}

type TokenService struct {
	config *config.Config
}

func NewTokenService(config *config.Config) *TokenService {
	return &TokenService{
		config: config,
	}
}
