package middleware

import (
	"github.com/gofiber/fiber/v3"
	"2Kang/pkg/utils"
)

func JWTProtected(c fiber.Ctx) error {
	_, err := utils.ValidateJWT(c)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "unauthorized"})
	}
	return c.Next()
}