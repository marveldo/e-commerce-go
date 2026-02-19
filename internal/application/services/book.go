package services

import (
	"strconv"

	"github.com/jinzhu/copier"
	"github.com/marveldo/gogin/internal/application/domain"
	"github.com/marveldo/gogin/internal/application/dto"
	"github.com/marveldo/gogin/internal/application/repository"
	"github.com/marveldo/gogin/internal/db"
)

type BookService struct {
	A *repository.AuthorRepository
	B *repository.BookRepository
}

func (s *BookService) CreateBook(bookInput *domain.BookInput) (*dto.BookResponsedto, error) {
	authors, err := s.A.CreateorGetAuthors(bookInput.Authors)
	if err != nil {
		return nil, err
	}
	bookInputWithoutAuthors := &domain.BookInputWithoutAuthors{}
	copier.Copy(bookInputWithoutAuthors, bookInput)
	bookModel, err := s.B.CreateBook(bookInputWithoutAuthors, authors)
	if err != nil {
		return nil, err
	}
	bookResponse := &dto.BookResponsedto{}
	copier.Copy(bookResponse, bookModel)
	return bookResponse, nil
}

func (s *BookService) FindAllBooks(query *domain.GetBookQuery) ([]*dto.BookResponsedto, error) {
	bookModels, err := s.B.FindAllBooks(query)
	if err != nil {
		return nil, err
	}
	bookResponses := []*dto.BookResponsedto{}
	for _, bookModel := range bookModels {
		bookResponse := &dto.BookResponsedto{}
		copier.Copy(bookResponse, bookModel)
		bookResponses = append(bookResponses, bookResponse)
	}
	return bookResponses, nil
}

func (s *BookService) DeleteBook(Id string) error {
	id, err := strconv.ParseUint(Id, 10, 32)
	if err != nil {
		return err
	}
	up_id := uint(id)
	query := &domain.GetBookQuery{
		ID: up_id,
	}
	return s.B.DeleteBook(query)
}

func (s *BookService) GetBookById(id string) (*db.Bookmodel, error) {
	uid, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return nil, err
	}
	return s.B.GetBookById(uint(uid))
}
