package repository

import (
	"errors"
	"fmt"

	"github.com/marveldo/gogin/internal/application/domain"
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

	order_item.Reference = fmt.Sprintf("ORD-%v", order_item.ID)
	err = tx.Save(order_item).Error
	if err != nil {
		return nil, err, nil, total
	}
	// err = tx.Preload("User").First(order_item , order_item.ID).Error
	// if err != nil {
	// 	return nil , err ,nil, total
	// }
	return order_item, nil, tx, total * 100
}

func (p *PaymentRepository) UpdatePaymentOrder(domain *domain.PaymentWebhookdomain) (*db.OrderModel, *gorm.DB, error) {
	tx := p.DB.Begin()
	var order_item db.OrderModel
	order_item.Reference = domain.Data.Reference
	err := tx.Where(&order_item).First(&order_item).Error
	if err != nil {
		return nil, nil, err
	}
	if order_item.IsProcessed {
		return nil ,nil, errors.New("Item already Proccesed")
	}
	switch domain.Event {
	case "charge.success":
		order_item.Status = db.Success
		order_item.IsProcessed = true
	case "charge.failed":
		order_item.Status = db.Failed
		order_item.IsProcessed = true
	default:
		order_item.Status = db.Pending
	}
	err = tx.Save(&order_item).Error
	if err != nil {
		return nil, nil, err
	}
	return &order_item, tx, nil

}
