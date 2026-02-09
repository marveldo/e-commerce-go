package db

import (
	"time"
)

type CartModel struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	UserID    uint     
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	// Relationships
	CartItems []CartItemModel `gorm:"foreignKey:CartID;constraint:OnDelete:CASCADE"`
	User      UserModel       `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
}

type CartItemModel struct {
	ID            uint `gorm:"primaryKey;autoIncrement"`
	CartID        uint 
	BookID        uint 
	Quantity      int  `gorm:"not null;default:1"`
	PriceSnapshot int  

	// Relationships
	Book Bookmodel `gorm:"foreignKey:BookID;references:ID"`
	Cart CartModel `gorm:"foreignKey:CartID;references:ID"`
}

// Unique constraint - Add this in migration or as a separate index
func (CartItemModel) TableName() string {
	return "cart_items"
}
