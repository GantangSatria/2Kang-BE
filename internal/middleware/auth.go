package middleware

import (
	"github.com/gofiber/fiber/v3"
	"2Kang/pkg/utils"
)

func JWTProtected(c fiber.Ctx) error {
	_, claims, err := utils.ValidateJWT(c)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "unauthorized"})
	}

	// Ambil ID dari claim
	idClaim := claims["id"]
	if idClaim == nil {
		return c.Status(401).JSON(fiber.Map{"error": "invalid token: no id"})
	}

	// Convert float64 → uint (JWT selalu float64)
	userIDFloat := idClaim.(float64)
	c.Locals("user_id", uint(userIDFloat))

	return c.Next()
}
