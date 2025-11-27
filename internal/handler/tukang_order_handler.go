package handler

import (
	"2Kang/internal/services"
	"github.com/gofiber/fiber/v3"
	"strconv"
)

type TukangOrderHandler struct {
	Service *services.TukangOrderService
}

func NewTukangOrderHandler(service *services.TukangOrderService) *TukangOrderHandler {
	return &TukangOrderHandler{Service: service}
}

func (h *TukangOrderHandler) Accept(c fiber.Ctx) error {
	orderID, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid order id"})
	}

	tukangIDStr := c.Locals("user_id").(string)
	tukangID, _ := strconv.ParseUint(tukangIDStr, 10, 32)

	err = h.Service.AcceptOrder(uint(orderID), uint(tukangID))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "order accepted"})
}

func (h *TukangOrderHandler) Start(c fiber.Ctx) error {
	orderID, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid order id"})
	}

	tukangIDStr := c.Locals("user_id").(string)
	tukangID, _ := strconv.ParseUint(tukangIDStr, 10, 32)

	err = h.Service.StartOrder(uint(orderID), uint(tukangID))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "order started"})
}

func (h *TukangOrderHandler) Finish(c fiber.Ctx) error {
	orderID, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid order id"})
	}

	tukangIDStr := c.Locals("user_id").(string)
	tukangID, _ := strconv.ParseUint(tukangIDStr, 10, 32)

	err = h.Service.FinishOrder(uint(orderID), uint(tukangID))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "order finished"})
}

func (h *TukangOrderHandler) Reject(c fiber.Ctx) error {
    orderIDParam := c.Params("id")
    orderID, _ := strconv.ParseUint(orderIDParam, 10, 32)
    tukangID := c.Locals("user_id").(uint)

    err := h.Service.RejectOrder(uint(orderID), tukangID)
    if err != nil {
        return c.Status(400).JSON(fiber.Map{"error": err.Error()})
    }

    return c.JSON(fiber.Map{"message": "Order berhasil dihapus (ditolak)"})
}
