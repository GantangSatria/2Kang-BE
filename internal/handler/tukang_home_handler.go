package handler

import (
	"2Kang/internal/services"
	"github.com/gofiber/fiber/v3"
	"strconv"
)

type TukangHomeHandler struct {
	TukangHomeService *services.TukangHomeService
}

func NewTukangHomeHandler(service *services.TukangHomeService) *TukangHomeHandler {
	return &TukangHomeHandler{
		TukangHomeService: service,
	}
}

func (h *TukangHomeHandler) GetHome(c fiber.Ctx) error {

	// Ambil user_id dari JWT token
	tukangIDStr := c.Locals("user_id")
	if tukangIDStr == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "unauthorized",
		})
	}

	idUint, err := strconv.ParseUint(tukangIDStr.(string), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid tukang ID",
		})
	}

	// Call service
	data, err := h.TukangHomeService.GetHome(uint(idUint))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "success",
		"data":    data,
	})
}
