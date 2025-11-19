package handler

import (
    "time"
    "strconv"
    "github.com/gofiber/fiber/v3"
    "2Kang/internal/services"
    "2Kang/pkg/dto/request"
    "2Kang/pkg/dto/response"
    "2Kang/internal/domain/entity"
)

type OrderHandler struct {
    TrxService *services.TransactionService
}

func NewOrderHandler(s *services.TransactionService) *OrderHandler {
    return &OrderHandler{TrxService: s}
}

// POST /api/v1/orders
func (h *OrderHandler) CreateOrder(c fiber.Ctx) error {
    // user must be authenticated
    uidI := c.Locals("user_id")
    if uidI == nil {
        return c.Status(401).JSON(fiber.Map{"error": "unauthorized"})
    }
    userID := uidI.(uint)

    var req request.CreateOrderRequest
    if err := c.Bind().Body(&req); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "invalid request"})
    }

    // parse scheduled time
    scheduledAt, err := time.Parse(time.RFC3339, req.ScheduledAt)
    if err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "invalid scheduled_at format, use ISO8601"})
    }

    // optional: server-side price validation (e.g. from tukang profile)
    // For now accept provided price but you can override by fetching tukang price.

    // build input map for service
    input := map[string]interface{}{
        "tukang_id":      float64(req.TukangID), // JSON->float64 when interface, keep consistent
        "scheduled_at":   scheduledAt,
        "address":        req.Address,
        "payment_method": req.PaymentMethod,
        "price":          req.Price,
        "notes":          req.Notes,
    }

    trx, err := h.TrxService.CreateOrder(userID, input)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": err.Error()})
    }

    res := response.OrderResponse{
        ID:            trx.ID,
        UserID:        trx.UserID,
        TukangID:      trx.TukangID,
        ScheduledAt:   trx.ScheduledAt,
        Address:       trx.Address,
        PaymentMethod: trx.PaymentMethod,
        Price:         trx.Price,
        Status:        string(trx.Status),
        Notes:         trx.Notes,
        CreatedAt:     trx.CreatedAt,
    }
    return c.Status(201).JSON(res)
}

// GET /api/v1/orders  -> returns orders of current authenticated user
func (h *OrderHandler) GetMyOrders(c fiber.Ctx) error {
    uidI := c.Locals("user_id")
    if uidI == nil {
        return c.Status(401).JSON(fiber.Map{"error": "unauthorized"})
    }
    userID := uidI.(uint)

    list, err := h.TrxService.GetOrdersByUser(userID)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": err.Error()})
    }

    resp := make([]response.OrderResponse, 0, len(list))
    for _, t := range list {
        resp = append(resp, response.OrderResponse{
            ID:            t.ID,
            UserID:        t.UserID,
            TukangID:      t.TukangID,
            ScheduledAt:   t.ScheduledAt,
            Address:       t.Address,
            PaymentMethod: t.PaymentMethod,
            Price:         t.Price,
            Status:        string(t.Status),
            Notes:         t.Notes,
            CreatedAt:     t.CreatedAt,
        })
    }
    return c.JSON(resp)
}

// GET /api/v1/orders/:id
func (h *OrderHandler) GetOrderDetail(c fiber.Ctx) error {
    idStr := c.Params("id")
    id64, err := strconv.ParseUint(idStr, 10, 64)
    if err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "invalid id"})
    }
    id := uint(id64)

    trx, err := h.TrxService.GetOrderByID(id)
    if err != nil {
        return c.Status(404).JSON(fiber.Map{"error": "order not found"})
    }
    res := response.OrderResponse{
        ID:            trx.ID,
        UserID:        trx.UserID,
        TukangID:      trx.TukangID,
        ScheduledAt:   trx.ScheduledAt,
        Address:       trx.Address,
        PaymentMethod: trx.PaymentMethod,
        Price:         trx.Price,
        Status:        string(trx.Status),
        Notes:         trx.Notes,
        CreatedAt:     trx.CreatedAt,
    }
    return c.JSON(res)
}

// TUKANG: get orders assigned to tukang (requires role check)
func (h *OrderHandler) GetOrdersForTukang(c fiber.Ctx) error {
    roleI := c.Locals("role")
    if roleI == nil || roleI.(string) != "tukang" {
        return c.Status(403).JSON(fiber.Map{"error": "forbidden"})
    }
    uidI := c.Locals("user_id")
    if uidI == nil {
        return c.Status(401).JSON(fiber.Map{"error": "unauthorized"})
    }
    tukangID := uidI.(uint)

    list, err := h.TrxService.GetOrdersByTukang(tukangID)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": err.Error()})
    }
    resp := make([]response.OrderResponse, 0, len(list))
    for _, t := range list {
        resp = append(resp, response.OrderResponse{
            ID:            t.ID,
            UserID:        t.UserID,
            TukangID:      t.TukangID,
            ScheduledAt:   t.ScheduledAt,
            Address:       t.Address,
            PaymentMethod: t.PaymentMethod,
            Price:         t.Price,
            Status:        string(t.Status),
            Notes:         t.Notes,
            CreatedAt:     t.CreatedAt,
        })
    }
    return c.JSON(resp)
}

// TUKANG: update order status (e.g. accept / start / done)
func (h *OrderHandler) UpdateOrderStatus(c fiber.Ctx) error {
    roleI := c.Locals("role")
    if roleI == nil || roleI.(string) != "tukang" {
        return c.Status(403).JSON(fiber.Map{"error": "forbidden"})
    }
    var body struct {
        OrderID uint   `json:"order_id"`
        Status  string `json:"status"` // expected: confirmed/ongoing/done/canceled
    }
    if err := c.Bind().Body(&body); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "invalid body"})
    }

    // validate status
    switch body.Status {
    case string(entity.StatusConfirmed), string(entity.StatusOngoing), string(entity.StatusDone), string(entity.StatusCanceled):
        // ok
    default:
        return c.Status(400).JSON(fiber.Map{"error": "invalid status"})
    }

    if err := h.TrxService.UpdateStatus(body.OrderID, entity.TransactionStatus(body.Status)); err != nil {
        return c.Status(500).JSON(fiber.Map{"error": err.Error()})
    }
    return c.JSON(fiber.Map{"message": "status updated"})
}
