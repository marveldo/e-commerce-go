package dto

type AuthorInputdto struct {
	Username string `json:"username" binding:"required"`
}
