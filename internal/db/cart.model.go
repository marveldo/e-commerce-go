package db

import (
	"time"
)

type CartModel struct {
	ID        uint            `gorm:"primaryKey;autoIncrement"`
	CreatedAt time.Time       `gorm:"autoCreateTime"`
	UpdatedAt time.Time       `gorm:"autoUpdateTime"`
	UserID    uint            `gorm:"uniqueIndex"`
	CartItems []CartItemModel `gorm:"foreignKey:CartID;constraint:OnDelete:CASCADE"`
}

type CartItemModel struct {
	ID            uint `gorm:"primaryKey;autoIncrement"`
	CartID        uint
	BookID        uint
	Quantity      int `gorm:"not null;default:1"`
	PriceSnapshot int
	Book          Bookmodel `gorm:"foreignKey:BookID;references:ID" json:"-"`
	Cart          CartModel `gorm:"foreignKey:CartID;references:ID" json:"-"`
}

func (CartItemModel) TableName() string {
	return "cart_items"
}
