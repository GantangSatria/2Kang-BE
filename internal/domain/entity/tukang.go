package entity
import "time"

type Tukang struct {
    ID        uint `gorm:"primaryKey;autoIncrement"`
    Name      string
    Email     string `gorm:"unique"`
    Password  string
    Category  string
    Bio       string
    Services  string
    Rating    float32
    CreatedAt time.Time
}
