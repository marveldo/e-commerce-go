package db

import "time"

type AuthorModel struct {
	ID        uint         `gorm:"primaryKey"`
	Username  string       `gorm:"unique"`
	Books     []*Bookmodel `gorm:"many2many:author_book"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
