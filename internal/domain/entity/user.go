package entity

import "time"

type Role string

const (
	RoleUser   Role = "user"
	RoleTukang Role = "tukang"
)

type User struct {
    ID        uint `gorm:"primaryKey;autoIncrement"`
    Name      string
    Email     string `gorm:"unique"`
    Password  string
    Role      Role // "user" / "tukang"
    Verified  bool
    CreatedAt time.Time
}
