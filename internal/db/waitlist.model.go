package db

type WaitlistModel struct {
	ID    uint   `gorm:"primaryKey"`
	User_ID uint `gorm:"uniqueIndex"`
	Books []Bookmodel `gorm:"many2many:waitlist_books"`
}