package db

import (
	"time"

	"gorm.io/gorm"
)

type OrderModel struct {
	ID        uint `gorm:"primaryKey; autoIncrement"`
	UserId    uint `gorm:"uniqueIndex"`
	Status    bool `gorm:"not null;default:false"`
	Price     int64
	Reference string `gorm:"unique"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (b *OrderModel) BeforeCreate(tx *gorm.DB) (err error) {
	b.Price = int64(b.Price * 100)
	return nil
}

func (b *OrderModel) AfterCreate(tx *gorm.DB) (err error) {
	b.Price = int64(b.Price / 100)
	return nil
}

func (b *OrderModel) AfterFind(tx *gorm.DB) (err error) {
	b.Price = int64(b.Price / 100)
	return nil
}
