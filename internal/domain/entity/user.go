package entity
import "time"

type User struct {
    ID        uint `gorm:"primaryKey;autoIncrement"`
    Name      string
    Email     string `gorm:"unique"`
    Password  string
    CreatedAt time.Time
}

