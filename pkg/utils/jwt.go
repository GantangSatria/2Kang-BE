package utils

import (
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
	"2Kang/internal/domain/entity"
)

var jwtSecret = []byte("SECRET_YANG_AMAN")

func GenerateJWT(user *entity.User) (string, error) {

	claims := jwt.MapClaims{
		"id":    user.ID,
		"email": user.Email,
		"role":  user.Role,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtSecret)
}

func ValidateJWT(c fiber.Ctx) (*jwt.Token, error) {

	tokenString := c.Get("Authorization")
	if tokenString == "" {
		return nil, fiber.ErrUnauthorized
	}

	return jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
}
