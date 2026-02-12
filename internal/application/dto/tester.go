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

type AdditionDto struct {
	Num_1 int `json:"num_1" binding:"required"`
	Num_2 int `json:"num_2" binding:"required"`
}