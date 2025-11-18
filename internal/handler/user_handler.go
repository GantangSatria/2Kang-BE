package handler

import (
	"github.com/gofiber/fiber/v3"
	"2Kang/internal/services"
)

type UserHandler struct {
	UserService services.UserService
}

func NewUserHandler(s services.UserService) *UserHandler {
	return &UserHandler{UserService: s}
}

func (h *UserHandler) GetProfile(c fiber.Ctx) error {
	userID := c.Locals("user_id").(string)

	profile, err := h.UserService.GetProfile(userID)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(profile)
}
