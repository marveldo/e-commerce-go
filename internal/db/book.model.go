package db

import "time"

type BookGenre string

const (
	Fantasy        BookGenre = "fantasy"
	ScienceFiction BookGenre = "scienceFiction"
	Mystery        BookGenre = "mystery"
	NonFiction     BookGenre = "nonFiction"
	Romance        BookGenre = "romance"
	Horror         BookGenre = "horror"
	THRILLER       BookGenre = "thriller"
	SELF_HELP      BookGenre = "self-help"
	HEALTH         BookGenre = "health"
	TRAVEL         BookGenre = "travel"
	CHILDREN       BookGenre = "children"
	ART            BookGenre = "art"
	COMICS         BookGenre = "comics"
	RELIGION       BookGenre = "religion"
	ACTION         BookGenre = "action"
	ADVENTURE      BookGenre = "adventure"
)

type Bookmodel struct {
	ID        uint `gorm:"primaryKey"`
	Title     string
	Owners    []*UserModel   `gorm:"many2many:owner_book"`
	Authors   []*AuthorModel `gorm:"many2many:author_book"`
	Genres    []BookGenre    `gorm:"type:jsonb;serializer:json"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
