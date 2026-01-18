package dto

type MessageInput struct {
	Message string `json:"message" binding:"required"`
}

type TestInput struct {
	Name    string `json:"name" binding:"required"`
	Message string `json:"message" binding:"required"`
}

type TestInputUpdate struct {
	Name    *string `json:"name" binding:"omitempty"`
	Message *string `json:"message" binding:"omitempty"`
}
