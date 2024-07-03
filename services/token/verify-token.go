package token_service

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

// VerifyToken verifies a token JWT validate
func (userTokenService *TokenService) VerifyToken(tokenString string) (jwt.MapClaims, error) {
	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("INVALID SIGNING METHOD")
		}

		return []byte(userTokenService.config.Auth.AccessTokenSecret), nil
	})

	// Check for errors
	if err != nil {
		return nil, err
	}

	// Validate the token
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("INVALID TOKEN")
}
