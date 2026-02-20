package dto

import "time"

type OrderDto struct {
	ID        uint      `json:"id"`
	UserId    uint      `json:"user_id"`
	Status    bool      `json:"status"`
	Price     int64     `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
type PaymentCreatedResponseDto struct {
	Order   OrderDto `json:"order"`
	AuthUrl string   `json:"auth_url"`
}

type CreatePaymentDto struct {
	CartID uint `json:"cart_id" binding:"required"`
}
