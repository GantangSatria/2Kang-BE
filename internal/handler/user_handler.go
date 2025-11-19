package handler

import (
	"2Kang/internal/services"
	"2Kang/pkg/dto/request"

	"github.com/gofiber/fiber/v3"
)

type UserHandler struct {
	UserService services.UserService
}

func NewUserHandler(s services.UserService) *UserHandler {
	return &UserHandler{UserService: s}
}

func (h *UserHandler) GetProfile(c fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	profile, err := h.UserService.GetProfile(userID)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(profile)
}

func (h *UserHandler) UpdateName(c fiber.Ctx) error {
    userID := c.Locals("user_id").(uint)

    var req request.UpdateUserNameRequest
    if err := c.Bind().Body(&req); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": err.Error()})
    }

    if err := h.UserService.UpdateName(userID, req.Name); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": err.Error()})
    }

    return c.JSON(fiber.Map{"message": "name updated"})
}
