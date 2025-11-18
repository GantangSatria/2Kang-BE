package handler

import (
	"github.com/gofiber/fiber/v3"
	"2Kang/internal/services"
	"2Kang/pkg/dto/request"
	"2Kang/pkg/dto/response"
)

type AuthHandler struct {
	service services.AuthService
}

func NewAuthHandler(service services.AuthService) *AuthHandler {
	return &AuthHandler{service}
}

func (h *AuthHandler) Register(c fiber.Ctx) error {
	var req request.RegisterRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	if err := h.service.Register(req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "registered successfully"})
}

func (h *AuthHandler) Login(c fiber.Ctx) error {
	var req request.LoginRequest
	_ = c.Bind().Body(&req)

	token, err := h.service.Login(req)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(response.AuthResponse{Token: token})
}
