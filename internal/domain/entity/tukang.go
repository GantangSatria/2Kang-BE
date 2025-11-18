package entity

type Tukang struct {
    ID        uint `gorm:"primaryKey;autoIncrement"`
    UserID    uint `gorm:"uniqueIndex"`
    Category  string
    Bio       string
    Services  string
    Rating    float32
}
