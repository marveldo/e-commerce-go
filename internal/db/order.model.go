package db

import (
	"time"

	"gorm.io/gorm"
)

type Status string

const (
	Pending Status = "PENDING"
	Success Status = "SUCCESS"
	Failed  Status = "FAILED"
)

type OrderModel struct {
	ID          uint   `gorm:"primaryKey; autoIncrement"`
	UserId      uint   `gorm:"index"`
	Status      Status `gorm:"not null;default:PENDING"`
	Price       int64
	Reference   string `gorm:"unique"`
	IsProcessed bool
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
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
