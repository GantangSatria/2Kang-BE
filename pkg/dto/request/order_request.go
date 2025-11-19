package request

type CreateOrderRequest struct {
    TukangID      uint    `json:"tukang_id"`                // id tukang yang dipesan (required)
    ScheduledAt   string  `json:"scheduled_at"`             // ISO8601 string, ex: "2025-11-30T09:00:00Z" (required)
    Address       string  `json:"address"`                  // lokasi layanan (required)
    PaymentMethod string  `json:"payment_method"`           // "cash" | "midtrans" | ... (required)
    Price         float64 `json:"price"`                    // harga yang disepakati (optional, server bisa override)
    Notes         string  `json:"notes,omitempty"`
}
