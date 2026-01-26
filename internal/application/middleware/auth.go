package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/marveldo/gogin/internal/application/utils"
	"github.com/marveldo/gogin/internal/config"
)

func Authmiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token_string := c.GetHeader("Authorization")
		if token_string == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":  http.StatusUnauthorized,
				"error": "Missing Authorization Token",
			})
			c.Abort()
			return
		}

		if len(token_string) > 7 && token_string[:7] == "Bearer " {
			token_string = token_string[7:]
		}

		token, err := jwt.ParseWithClaims(token_string, &utils.Claims{}, func(j *jwt.Token) (interface{}, error) {
			return []byte(config.LoadConfig().JWTSecret), nil
		})
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":  http.StatusUnauthorized,
				"error": "Invalid token",
			})
			c.Abort()
			return
		}
		if claims, ok := token.Claims.(*utils.Claims); ok {
			if claims.Type != utils.ACCESS {
				c.JSON(http.StatusUnauthorized, gin.H{
					"code":  http.StatusUnauthorized,
					"error": "Invalid token Type",
				})
				c.Abort()
			} else {
				c.Set("username", claims.Username)
				c.Set("user_id", claims.Id)
				c.Next()
			}

		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":  http.StatusUnauthorized,
				"error": "Invalid token claims",
			})
			c.Abort()
		}
	}
}
