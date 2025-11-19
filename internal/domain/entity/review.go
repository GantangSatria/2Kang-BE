package entity

import "time"

type Review struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	UserID    uint      `json:"user_id"`
	TukangID  uint      `json:"tukang_id"`
	Rating    float32   `json:"rating"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
}
