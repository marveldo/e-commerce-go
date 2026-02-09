package services

import (
	"strconv"

	"github.com/marveldo/gogin/internal/application/domain"
	"github.com/marveldo/gogin/internal/application/repository"
	"github.com/marveldo/gogin/internal/db"
)

type AuthorService struct {
	R *repository.AuthorRepository
}

func (s *AuthorService) CreateAuthor(author *domain.AuthorInput) (*db.AuthorModel, error) {
	return s.R.CreateAuthor(author)
}

func (s *AuthorService) DeleteAuthor(Id string) error {
	id, err := strconv.ParseUint(Id, 10, 32)
	if err != nil {
		return err
	}
	up_id := uint(id)
	query := &domain.GetAuthorQuery{
		ID: &up_id,
	}

	return s.R.DeleteAuthor(query)
}

func (s *AuthorService) GetallAllAuthors() ([]*db.AuthorModel, error) {
	return s.R.GetallAuthors()
}

