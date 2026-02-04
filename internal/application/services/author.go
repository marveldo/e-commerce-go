package services

import (
	"github.com/marveldo/gogin/internal/application/domain"
	"github.com/marveldo/gogin/internal/application/repository"
	"github.com/marveldo/gogin/internal/db"
)

type AuthorService struct {
	R *repository.AuthorRepository
	U *repository.Userrespository
}

func (s *AuthorService) CreateAuthor(authorInput *domain.AuthorInput , user_id uint) (*db.UserModel, error){
	query := &domain.GetUserQuery{
		ID: &user_id,
	}
	user , err := s.U.GetUser(query)
	if err != nil {
		return nil, err
	}
	return s.R.CreateAuthor(user , authorInput)
}
