package services

import (
	"github.com/marveldo/gogin/internal/application/repository"
	"github.com/marveldo/gogin/internal/db"
)

type WaitlistService struct {
	W *repository.WaitlistRepository
	U *repository.Userrespository
	B *repository.BookRepository
}

func (s *WaitlistService) GetBooksInWaitlist(userId uint) ([]db.Bookmodel, error) {
	return s.W.GetBooksInWaitlist(userId)
}

func (s *WaitlistService) AddBookToWaitlist(userId uint, bookId uint) (*db.Bookmodel, error) {
	book, err := s.B.GetBookById(bookId)
	if err != nil {
		return nil, err
	}
	return s.W.AddBookToWaitlist(userId, book)
}