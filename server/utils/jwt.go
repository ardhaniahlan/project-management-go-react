package utils

import (
	"project-management-be/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func GenerateTokenJWT(userID uint64, role string, email string, publicID uuid.UUID) (string, error) {
	secret := config.AppConfig.JWTSecret
	exp, err := time.ParseDuration(config.AppConfig.JWTExpired)
	if err != nil {
		return "", err
	}
	expired := time.Now().Add(exp)
	
	claims := jwt.MapClaims{
		"user_id": userID,
		"role": role,
		"email": email,
		"public_id": publicID,
		"exp": expired.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func GenerateRefreshTokenJWT(userID uint64) (string, error) {
	secret := config.AppConfig.JWTSecret
	exp, err := time.ParseDuration(config.AppConfig.JWTRefreshToken)
	if err != nil {
		return "", err
	}
	expired := time.Now().Add(exp)
	
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp": expired.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}