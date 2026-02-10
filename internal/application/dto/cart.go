package dto

type CartItemInputDto struct {
	BookID   uint `json:"book_id" binding:"required"`
	Quantity int  `json:"quantity" binding:"required,gt=0"`
}
