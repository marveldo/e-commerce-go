package repository

import (
	"github.com/jinzhu/copier"
	"github.com/marveldo/gogin/internal/application/domain"
	"github.com/marveldo/gogin/internal/db"
	"gorm.io/gorm"
)

type BookRepository struct {
	DB *gorm.DB
}

func (r *BookRepository) CreateBook(book *domain.BookInputWithoutAuthors, authors []*db.AuthorModel) (*db.Bookmodel, error) {
	bookModel := db.Bookmodel{}
	copier.Copy(&bookModel, book)
	err := r.DB.Create(&bookModel).Error
	if err != nil {
		return nil, err
	}
	err = r.DB.Model(&bookModel).Association("Authors").Append(authors)
	if err != nil {
		return nil, err
	}
	err = r.DB.Preload("Authors").First(&bookModel, bookModel.ID).Error
	if err != nil {
		return nil, err
	}
	return &bookModel, nil
}

func (r *BookRepository) FindAllBooks() ([]*db.Bookmodel, error) {
	var books []*db.Bookmodel
	err := r.DB.Preload("Authors").Find(&books).Error
	return books, err
}

func (r *BookRepository) DeleteBook(book *domain.GetBookQuery) error {
	bookModel := &db.Bookmodel{}
	err := copier.Copy(bookModel, book)
	if err != nil {
		return err
	}
	err = r.DB.Model(bookModel).Association("Authors").Clear()
	if err != nil {
		return err
	}
	err = r.DB.Delete(bookModel).Error
	if err != nil {
		return err
	}
	return nil
}
