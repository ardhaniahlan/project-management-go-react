package middleware

import (
	"strings"
	"project-management-be/config"
	"project-management-be/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func Protected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return utils.Unauthorized(c, "Akses ditolak", "Token tidak ditemukan")
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return utils.Unauthorized(c, "Format token salah", "Gunakan format: Bearer <token>")
		}

		tokenString := parts[1]
		secret := []byte(config.AppConfig.JWTSecret)

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fiber.ErrUnauthorized
			}
			return secret, nil
		})

		if err != nil || !token.Valid {
			return utils.Unauthorized(c, "Akses ditolak", "Token tidak valid atau sudah kadaluarsa")
		}

		claims := token.Claims.(jwt.MapClaims)
		c.Locals("user", claims)

		return c.Next()
	}
}