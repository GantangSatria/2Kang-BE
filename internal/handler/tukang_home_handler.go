package handler

import (
	"2Kang/internal/services"
	"log"
	// "errors"
	"fmt"

	"github.com/gofiber/fiber/v3"
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
    log.Println("ROLE:", c.Locals("role"))
    log.Println("USER_ID:", c.Locals("user_id"))
    tukangID := c.Locals("user_id").(uint)  // dari middleware JWT



    // Validasi: cek apakah tukang dengan ID ini ada
    //tukang, err := h.TukangHomeService.TukangRepo.FindByID(tukangID)
    // if err != nil {
    //     fmt.Printf("DEBUG: Tukang not found with ID %d: %v\n", tukangID, err)
    //     return c.Status(400).JSON(fiber.Map{"error": "ID tidak valid"})
    // }
    
    //fmt.Printf("DEBUG: Found tukang: %s (ID: %d)\n", tukang.Name, tukang.ID)
    
    // Get home data
    res, err := h.TukangHomeService.GetHome(tukangID)
    if err != nil {
        fmt.Printf("DEBUG: GetHome error: %v\n", err)
        return c.Status(500).JSON(fiber.Map{"error": err.Error()})
    }
    
    return c.JSON(res)
}