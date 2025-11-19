package response

import "time"

type OrderResponse struct {
    ID            uint      `json:"id"`
    UserID        uint      `json:"user_id"`
    TukangID      uint      `json:"tukang_id"`
    ScheduledAt   time.Time `json:"scheduled_at"`
    Address       string    `json:"address"`
    PaymentMethod string    `json:"payment_method"`
    Price         float64   `json:"price"`
    Status        string    `json:"status"`
    Notes         string    `json:"notes"`
    CreatedAt     time.Time `json:"created_at"`
}
