package utils

import (
	"os"
	"strings"
	"time"

	// "2Kang/internal/domain/entity"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func GenerateJWT(id uint, email string, role string) (string, error) {
	claims := jwt.MapClaims{
		"id":    id,
		"email": email,
		"role":  role,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}


func ValidateJWT(c fiber.Ctx) (*jwt.Token, jwt.MapClaims, error) {
	auth := c.Get("Authorization")
	if auth == "" {
		return nil, nil, fiber.ErrUnauthorized
	}

	parts := strings.Split(auth, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return nil, nil, fiber.ErrUnauthorized
	}

	tokenStr := parts[1]

	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, nil, err
	}

	claims := token.Claims.(jwt.MapClaims)

	return token, claims, nil
}


// func JwtSecret() []byte {
//     return []byte(os.Getenv("JWT_SECRET"))
// }