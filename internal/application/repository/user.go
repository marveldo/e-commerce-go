package repository

import (
	"github.com/jinzhu/copier"
	"github.com/marveldo/gogin/internal/application/domain"
	"github.com/marveldo/gogin/internal/db"
	"gorm.io/gorm"
)

type Userrespository struct {
	DB *gorm.DB
}

func (u *Userrespository) Createuser(i *domain.UserInput) (*db.UserModel, error) {
	empty_cart := db.CartModel{
		CartItems: []db.CartItemModel{},
	}
	empty_waitlist := db.WaitlistModel{
		Books: []db.Bookmodel{},
	}
	user := db.UserModel{
		Username: i.Username,
		Email:    i.Email,
		Bio:      i.Bio,
		Password: i.Password,
		Cart:     empty_cart,
		Waitlist: empty_waitlist,
	}
	err := u.DB.Create(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *Userrespository) GetUser(i *domain.GetUserQuery) (*db.UserModel, error) {
	user := &db.UserModel{}
	err := copier.Copy(user, i)
	if err != nil {
		return nil, err
	}
	result := u.DB.Preload("Cart").Preload("Waitlist").Preload("Books").Where(user).First(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (u *Userrespository) GetUserWithtx(tx *gorm.DB, i *domain.GetUserQuery) (*db.UserModel, error) {
	user := &db.UserModel{}
	err := copier.Copy(user, i)
	if err != nil {
		return nil, err
	}
	result := tx.Preload("Cart").Preload("Waitlist").Where(user).First(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}
