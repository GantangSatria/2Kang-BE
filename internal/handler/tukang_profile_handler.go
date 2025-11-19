package handler

import (
	"2Kang/internal/services"
	"2Kang/pkg/dto/request"

	"github.com/gofiber/fiber/v3"
)

type TukangProfileHandler struct {
	Service *services.TukangProfileService
}

func NewTukangProfileHandler(service *services.TukangProfileService) *TukangProfileHandler {
	return &TukangProfileHandler{Service: service}
}

func (h *TukangProfileHandler) GetProfile(c fiber.Ctx) error {
	tukangID := c.Locals("user_id").(uint)

	res, err := h.Service.GetProfile(tukangID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(res)
}

func (h *TukangProfileHandler) UpdateProfile(c fiber.Ctx) error {

	tukangID := c.Locals("user_id").(uint)

	var req request.UpdateTukangProfileRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid request"})
	}

	if err := h.Service.UpdateProfile(tukangID, req); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "profile updated"})
}