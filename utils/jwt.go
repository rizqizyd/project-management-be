package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/rizqizyd/project-management-be/config"
)

func GenerateToken(userID int64, role, email string, publicID uuid.UUID) (string, error) {
	secret := config.AppConfig.JWTSecret
	duration, _ := time.ParseDuration(config.AppConfig.JWTExpire)

	claims := jwt.MapClaims{
		"user_id":   userID,
		"role":      role,
		"public_id": publicID,
		"email":     email,
		"exp":       time.Now().Add(duration).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	return token.SignedString([]byte(secret))
}
