package usecase

import (
	"golang-boilerplate/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func (tokenUsecase *UserTokenUsecase) CreateRefreshToken(user *models.User) (t string, err error) {
	claimsRefresh := &JwtCustomRefreshClaims{
		ID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 8)),
		},
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsRefresh)

	rt, err := refreshToken.SignedString([]byte(tokenUsecase.config.Auth.RefreshTokenSecret))
	if err != nil {
		return "", err
	}
	return rt, err
}
