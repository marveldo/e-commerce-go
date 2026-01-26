package domain

import "github.com/marveldo/gogin/internal/db"

type UserInput struct {
	Email    string
	Username string
	Bio      *string
	Password string
}

type GetUserQuery struct {
	Email    *string
	Username *string
	ID       *uint
}

type LoginInput struct {
	Email    string
	Password string
}

type LoginResponse struct {
	User    *db.UserModel `json:"user"`
	Access  string   `json:"access"`
	Refresh string   `json:"refresh"`
}
