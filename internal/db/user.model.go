package db

import "time"

type UserRole string

const (
	ADMIN UserRole = "admin"
	USER  UserRole = "user"
)

type UserModel struct {
	ID        uint   `gorm:"primaryKey"`
	Username  string `gorm:"unique"`
	Email     string `gorm:"unique"`
	Bio       *string
	Password  string   `json:"-"`
	Role      UserRole `gorm:"type:text;default:user"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
