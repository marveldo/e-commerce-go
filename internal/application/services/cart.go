package services

import (
	"github.com/marveldo/gogin/internal/application/domain"
	"github.com/marveldo/gogin/internal/application/repository"
	"github.com/marveldo/gogin/internal/db"
)

type CartService struct {
	R *repository.CartRepository
	U *repository.Userrespository
}


func (s *CartService) GetCartItems(user_id uint) ([]db.CartItemModel , error) {
	user , err := s.U.GetUser(&domain.GetUserQuery{ID: &user_id})
	if err != nil {
		return nil , err
	}
	return s.R.GetCartItems(user.Cart.ID)
}

func (s *CartService) AddCartItem(user_id uint, input *domain.CartItemInputDomain) (*db.CartItemModel, error) {
	user , err := s.U.GetUser(&domain.GetUserQuery{ID: &user_id})
	if err != nil {
		return nil , err
	}
	return s.R.AddCartItem(user.Cart.ID, input)
}
