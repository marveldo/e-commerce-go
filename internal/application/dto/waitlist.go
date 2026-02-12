package dto

type AddBooktoWaitlist struct {
	Book_id uint `json:"book_id" binding:"required"`
}

