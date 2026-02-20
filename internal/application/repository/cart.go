package repository

import (
	"github.com/jinzhu/copier"
	"github.com/marveldo/gogin/internal/application/domain"
	"github.com/marveldo/gogin/internal/db"
	"gorm.io/gorm"
)

type CartRepository struct {
	DB *gorm.DB
}

func (r *CartRepository) GetCartItems(cartId uint) ([]db.CartItemModel, error) {
	var cartItems []db.CartItemModel
	err := r.DB.Preload("Book").Where("cart_id = ?", cartId).Find(&cartItems).Error
	return cartItems, err
}

func (r *CartRepository) GetCartItemsWithtx(tx *gorm.DB, cartId uint) ([]db.CartItemModel, error) {
	var cartItems []db.CartItemModel
	err := tx.Preload("Book").Where("cart_id = ?", cartId).Find(&cartItems).Error
	return cartItems, err
}

func (r *CartRepository) AddCartItem(cart_id uint, input *domain.CartItemInputDomain) (*db.CartItemModel, error) {
	cartitem := &db.CartItemModel{}
	cart := &db.CartModel{}
	result := r.DB.Model(&db.CartModel{}).Where("id = ?", cart_id).First(cart)
	if result.Error != nil {
		return nil, result.Error
	}
	result = r.DB.Preload("Book").Where("cart_id = ?", cart_id).Where("book_id = ?", input.BookID).First(cartitem)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			copier.Copy(cartitem, input)
			cartitem.CartID = cart_id
			err := r.DB.Create(cartitem).Error
			err = r.DB.Preload("Book").First(cartitem, cartitem.ID).Error
			return cartitem, err
		}
		return nil, result.Error
	}
	cartitem.Quantity += input.Quantity
	err := r.DB.Save(cartitem).Error
	if err != nil {
		return nil, err
	}
	return cartitem, nil

}

// func (r *CartRepository) RemoveCartItem(cart_id uint, book_id uint) error {
// 	result := r.DB.Where("cart_id = ?", cart_id).Where("book_id = ?", book_id).Delete(&db.CartItemModel{})
// 	return result.Error
// }

// func (r *CartRepository) ClearCart(cart_id uint) error {
