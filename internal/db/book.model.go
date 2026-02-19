package db

import (
	"time"

	"gorm.io/gorm"
)

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
	ID          uint `gorm:"primaryKey"`
	Title       string
	Price       int64
	Description string
	Owners      []*UserModel    `gorm:"many2many:owner_book"`
	Authors     []*AuthorModel  `gorm:"many2many:author_book"`
	Genres      []BookGenre     `gorm:"type:jsonb;serializer:json"`
	Waitlists   []WaitlistModel `gorm:"many2many:waitlist_books"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (b *Bookmodel) BeforeCreate(tx *gorm.DB) (err error) {
	b.Price = int64(b.Price * 100)
	return nil
}

func (b *Bookmodel) AfterCreate(tx *gorm.DB) (err error) {
	b.Price = int64(b.Price / 100)
	return nil
}

func (b *Bookmodel) AfterFind(tx *gorm.DB) (err error) {
	b.Price = int64(b.Price / 100)
	return nil
}
