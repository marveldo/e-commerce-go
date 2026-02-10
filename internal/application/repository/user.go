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
	user := db.UserModel{
		Username: i.Username,
		Email:    i.Email,
		Bio:      i.Bio,
		Password: i.Password,
		Cart:     empty_cart,
	}
	err := u.DB.Create(&user).Error
	return &user, err
}

func (u *Userrespository) GetUser(i *domain.GetUserQuery) (*db.UserModel, error) {
	user := &db.UserModel{}
	err := copier.Copy(user, i)
	if err != nil {
		return nil, err
	}
	result := u.DB.Preload("Cart").Where(user).First(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}
