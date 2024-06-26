package token_service

import (
	"golang-boilerplate/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func (userTokenService *UserTokenService) CreateAccessToken(user *models.User) (t string, expired int64, err error) {
	exp := time.Now().Add(time.Hour * time.Duration(8))
	claims := &JwtCustomClaims{
		user.Username,
		user.ID,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}
	expired = exp.Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err = token.SignedString([]byte(userTokenService.config.Auth.AccessTokenSecret))
	if err != nil {
		return
	}

	return
}
