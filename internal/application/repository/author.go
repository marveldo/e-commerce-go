package repository

import (
	"github.com/marveldo/gogin/internal/application/domain"
	"github.com/marveldo/gogin/internal/db"
	"gorm.io/gorm"
)

type AuthorRepository struct {
	DB *gorm.DB
}

func (a *AuthorRepository) CreateAuthor(user *db.UserModel, author *domain.AuthorInput) (*db.UserModel, error) {
	authorM := db.AuthorModel{
		Username: author.Username,
	}

	var existingAuthor db.AuthorModel
	err := a.DB.Where("user_id = ?", user.ID).First(&existingAuthor).Error
	if err == nil {
		return nil, gorm.ErrForeignKeyViolated
	}
	err = a.DB.Model(user).Association("Author").Append(&authorM)
	if err != nil {
		return nil, err
	}
	result := a.DB.Preload("Author").First(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil

}
