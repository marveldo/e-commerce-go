package db

import "time"

type AuthorModel struct {
	ID        uint         `gorm:"primaryKey"`
	Username  string       `gorm:"unique"`
	Book      []*Bookmodel `gorm:"many2many:author_book"`
	UserID    *uint
	CreatedAt time.Time
	UpdatedAt time.Time
}
