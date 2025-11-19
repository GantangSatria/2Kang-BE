package entity

import "time"

type TransactionStatus string

const (
	StatusPending   TransactionStatus = "pending"   // baru dibuat, belum dibayar/konfirmasi
	StatusConfirmed TransactionStatus = "confirmed" // user sudah konfirmasi / pembayaran sukses
	StatusOngoing   TransactionStatus = "ongoing"   // sedang dikerjakan tukang
	StatusDone      TransactionStatus = "done"      // selesai
	StatusCanceled  TransactionStatus = "canceled"  // batal
)

type Transaction struct {
	ID            uint              `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID        uint              `json:"user_id"`   // pemesan
	TukangID      uint              `json:"tukang_id"` // tukang yang dipesan
	ScheduledAt   time.Time         `json:"scheduled_at"`
	Address       string            `json:"address"`
	PaymentMethod string            `json:"payment_method"` // e.g. "cash", "midtrans", "gopay"
	Price         float64           `json:"price"`
	Status        TransactionStatus `gorm:"type:varchar(20)" json:"status"`
	Notes         string            `json:"notes"`
	CreatedAt     time.Time         `json:"created_at"`
	UpdatedAt     time.Time         `json:"updated_at"`
}
