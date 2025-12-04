package utils

import (
	"fmt"
	"redirect-service/config"

	"github.com/golang-jwt/jwt"
)

func ValidateJWT(tokenString string) (jwt.MapClaims, error) {
	cfg := config.GetConfig()
	jwtKey := []byte(cfg.JwtSecret)

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("wrong authentication method")
		}

		return jwtKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("error jwt token parse: %w", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("wronk token")
}
