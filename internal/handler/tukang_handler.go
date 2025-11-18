package handler

import (
	"github.com/gofiber/fiber/v3"
	"2Kang/internal/services"
)

type TukangHandler struct {
	TukangService services.TukangService
}

func NewTukangHandler(s services.TukangService) *TukangHandler {
	return &TukangHandler{TukangService: s}
}

func (h *TukangHandler) GetTukangList(c fiber.Ctx) error {
	kategori := c.Query("kategori")

	result, err := h.TukangService.GetTukangList(kategori)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(result)
}

func (h *TukangHandler) GetTukangDetail(c fiber.Ctx) error {
	userID := c.Params("user_id")

	result, err := h.TukangService.GetTukangDetail(userID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(result)
}
