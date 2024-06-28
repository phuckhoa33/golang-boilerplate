package token_service

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

func (userTokenService *TokenService) CreateFogotPasswordToken(issuer string, subject string, duration time.Duration) (string, error) {
	// Define the signing method and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": issuer,                          // Issuer
		"sub": subject,                         // Subject
		"exp": time.Now().Add(duration).Unix(), // Expiration time
	})

	// Secret key used to sign the token
	// IMPORTANT: Keep this safe and do not hard-code in production!
	secretKey := []byte("your_secret_key")

	// Sign the token with the secret key
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
