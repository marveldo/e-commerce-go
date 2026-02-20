package repository

import (
	"github.com/marveldo/gogin/internal/db"
	"gorm.io/gorm"
)

type PaymentRepository struct {
	DB *gorm.DB
}

func (p *PaymentRepository) CreatePaymentOrder(cart_items []db.CartItemModel, user_id uint) (*db.OrderModel, error, *gorm.DB, float64) {
	var total float64
	for _, item := range cart_items {
		total += float64(item.Book.Price) * float64(item.Quantity)
	}
	order_item := &db.OrderModel{
		UserId: user_id,
		Price:  int64(total),
	}
	tx := p.DB.Begin()
	if tx.Error != nil {
		return nil, tx.Error, nil, total
	}
	err := tx.Create(order_item).Error
	if err != nil {
		return nil, err, nil, total
	}
	// err = tx.Preload("User").First(order_item , order_item.ID).Error
	// if err != nil {
	// 	return nil , err ,nil, total
	// }
	return order_item, nil, tx, total * 100
}
