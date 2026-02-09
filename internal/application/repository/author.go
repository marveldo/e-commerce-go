package repository

import (
	"github.com/jinzhu/copier"
	"github.com/marveldo/gogin/internal/application/domain"
	"github.com/marveldo/gogin/internal/db"
	"gorm.io/gorm"
)

type AuthorRepository struct {
	DB *gorm.DB
}

func (r *AuthorRepository) CreateorGetAuthors(authors []domain.AuthorInput) ([]*db.AuthorModel, error) {
	var authorModels []*db.AuthorModel
	for _, authorInput := range authors {
		var authorModel db.AuthorModel
		result := r.DB.Where(db.AuthorModel{Username: authorInput.Username}).FirstOrCreate(&authorModel)
		if result.Error != nil {
			return nil, result.Error
		}
		authorModels = append(authorModels, &authorModel)
	}
	return authorModels, nil
}

func (r *AuthorRepository) GetallAuthors() ([]*db.AuthorModel, error) {
	var authors []*db.AuthorModel
	err := r.DB.Preload("Books").Find(&authors).Error
	return authors, err
}

func (r *AuthorRepository) CreateAuthor(author *domain.AuthorInput) (*db.AuthorModel, error) {
	authorModel := &db.AuthorModel{}
	err := copier.Copy(authorModel, author)
	if err != nil {
		return nil, err
	}
	err = r.DB.Create(authorModel).Error
	return authorModel, err
}

func (r *AuthorRepository) DeleteAuthor(author *domain.GetAuthorQuery) error {
	authorModel := &db.AuthorModel{}
	err := copier.Copy(authorModel, author)
	if err != nil {
		return err
	}
	err = r.DB.Delete(authorModel).Error
	if err != nil {
		return err
	}
	return nil
}
