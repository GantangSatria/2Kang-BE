package entity

import "time"

type User struct {
    ID        string `gorm:"primaryKey"`
    Name      string
    Email     string `gorm:"unique"`
    Password  string
    Role      string // "user" / "tukang"
    Verified  bool
    CreatedAt time.Time
}
