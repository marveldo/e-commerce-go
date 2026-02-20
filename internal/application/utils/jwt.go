package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/marveldo/gogin/internal/config"
)

type TokenType string

const (
	ACCESS  TokenType = "access"
	REFRESH TokenType = "refresh"
)

type Claims struct {
	Username string    `json:"username"`
	Id       uint      `json:"id"`
	Type     TokenType `json:"type"`
	jwt.RegisteredClaims
}

var jwt_secret = func() []byte {
	return []byte(config.LoadConfig().JWTSecret)
}

func GenrateJwtToken(un string, id uint, ty string, exp_hour int) (string, error) {
	if ty != "access" && ty != "refresh" {
		return "", errors.New("Token Type Must be Access or Refresh")
	}
	claims := Claims{
		Username: un,
		Id:       id,
		Type:     TokenType(ty),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(exp_hour) * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwt_secret())
}
