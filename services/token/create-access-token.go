package token_service

import (
	"golang-boilerplate/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func (userTokenService *TokenService) CreateAccessToken(user *models.User) (t string, expired int64, err error) {
	accessTokenExpiredIn, _ := time.ParseDuration(userTokenService.config.Auth.AccessTokenExpiredIn)
	claims := &JwtCustomClaims{
		user.Username,
		user.ID,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(accessTokenExpiredIn)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err = token.SignedString([]byte(userTokenService.config.Auth.AccessTokenSecret))
	if err != nil {
		return
	}

	return
}
