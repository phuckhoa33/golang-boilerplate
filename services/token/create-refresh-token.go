package token_service

import (
	"golang-boilerplate/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func (userTokenService *TokenService) CreateRefreshToken(user *models.User) (t string, err error) {
	refreshTokenExpiredIn, _ := time.ParseDuration(userTokenService.config.Auth.RefreshTokenExpiredIn)
	claimsRefresh := &JwtCustomRefreshClaims{
		UserID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(refreshTokenExpiredIn)),
		},
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsRefresh)

	rt, err := refreshToken.SignedString([]byte(userTokenService.config.Auth.RefreshTokenSecret))
	if err != nil {
		return "", err
	}
	return rt, err
}
