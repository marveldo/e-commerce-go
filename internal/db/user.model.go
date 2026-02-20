package db

import "time"

type UserRole string

const (
	ADMIN UserRole = "admin"
	USER  UserRole = "user"
)

type UserModel struct {
	ID        uint   `gorm:"primaryKey"`
	Username  string `gorm:"unique"`
	Email     string `gorm:"unique"`
	Bio       *string
	Password  string        `json:"-"`
	Books     []*Bookmodel  `gorm:"many2many:owner_book"`
	Waitlist  WaitlistModel `gorm:"foreignKey:User_ID;constraint:OnDelete:CASCADE"`
	Cart      CartModel     `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Role      UserRole      `gorm:"type:text;default:user"`
	Orders    []OrderModel  `gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
