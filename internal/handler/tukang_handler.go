package handler

import (
	"2Kang/internal/services"
	"2Kang/pkg/dto/request"
	"strconv"

	"github.com/gofiber/fiber/v3"
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

	paramID := c.Params("id") 

    id, err := strconv.ParseUint(paramID, 10, 32)
    if err != nil {
        return c.Status(400).JSON(fiber.Map{
            "error": "ID tidak valid",
        })
    }
	
    result, err := h.TukangService.GetTukangDetail(uint(id))
    if err != nil {
        return c.Status(404).JSON(fiber.Map{
            "error": "Tukang tidak ditemukan",
        })
    }

	return c.JSON(result)
}

func (h *TukangHandler) UpdateCategory(c fiber.Ctx) error {
    tukangID := c.Locals("user_id").(uint)

    var req request.UpdateTukangCategoryRequest
    if err := c.Bind().Body(&req); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": err.Error()})
    }

    if err := h.TukangService.UpdateCategory(tukangID, req.Category); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": err.Error()})
    }

    return c.JSON(fiber.Map{"message": "category updated"})
}

func (h *TukangHandler) UpdateBio(c fiber.Ctx) error {
    tukangID := c.Locals("user_id").(uint)

    var req request.UpdateTukangBioRequest
    if err := c.Bind().Body(&req); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": err.Error()})
    }

    if err := h.TukangService.UpdateBio(tukangID, req.Bio); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": err.Error()})
    }

    return c.JSON(fiber.Map{"message": "bio updated"})
}

func (h *TukangHandler) UpdateServices(c fiber.Ctx) error {
    tukangID := c.Locals("user_id").(uint)

    var req request.UpdateTukangServicesRequest
    if err := c.Bind().Body(&req); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": err.Error()})
    }

    if err := h.TukangService.UpdateServices(tukangID, req.Services); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": err.Error()})
    }

    return c.JSON(fiber.Map{"message": "services updated"})
}
