package auth

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// JWT https://developers.liquid.com/#authentication
func JWT(path string, tokenID, secret string) (string, error) {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"path":     path,
		"nonce":    time.Now().UnixNano(),
		"token_id": tokenID,
	})

	signed, err := jwtToken.SignedString([]byte(secret))
	if err != nil {
		return "", fmt.Errorf("can't sign jwt: %w", err)
	}

	return signed, nil
}
