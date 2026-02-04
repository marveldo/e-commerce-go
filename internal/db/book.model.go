package db
import "time"

type BookGenre struct {
	ID   uint `gorm:"primaryKey"`
	Name string
	Book []*Bookmodel `gorm:"many2many:bookgenre_book"`
}
type Bookmodel struct {
	ID     uint `gorm:"primaryKey"`
	Title  string
	Owner  [] *UserModel `gorm:"many2many:owner_book"`
	Author [] *AuthorModel `gorm:"many2many:author_book"`
	Genre  []*BookGenre `gorm:"many2many:bookgenre_book"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
