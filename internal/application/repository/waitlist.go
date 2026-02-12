package repository

import (
	"github.com/marveldo/gogin/internal/db"
	"gorm.io/gorm"
)

type WaitlistRepository struct {
	DB *gorm.DB
}

func (r *WaitlistRepository) GetBooksInWaitlist(userId uint) ([]db.Bookmodel, error) {
	var waitlist db.WaitlistModel
	err := r.DB.Preload("Books").Where("user_id = ?", userId).First(&waitlist).Error
	if err != nil {
		return nil, err
	}
	return waitlist.Books, nil
}

func (r *WaitlistRepository) AddBookToWaitlist(userId uint, book *db.Bookmodel) (*db.Bookmodel, error) {
	var waitlist db.WaitlistModel
	err := r.DB.Preload("Books").Where("user_id = ?", userId).First(&waitlist).Error
	if err != nil {
		return nil , err
	}
	err = r.DB.Model(&waitlist).Association("Books").Append(book)
	return book, err
}
