package utils

import (
	"auth-service/model"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateJWT(user model.User) (string, error) {
	jwtKey := []byte("k1lroyVoiceChat")
	claims := jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", fmt.Errorf("ошибка в подписи jwt токена: %w", err)
	}

	return tokenString, nil
}
