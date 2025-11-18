package entity

type Tukang struct {
    ID        string `gorm:"primaryKey"`
    UserID    string `gorm:"uniqueIndex"`
    Category  string
    Bio       string
    Services  string
    Rating    float32
}
