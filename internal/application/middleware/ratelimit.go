package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/marveldo/gogin/internal/config"
	"golang.org/x/time/rate"
)

func RatelimitMiddleware(config *config.Config) gin.HandlerFunc {
	limiter := rate.NewLimiter(rate.Limit(config.Request_Per_Second), config.Request_Burst)
	return func(c *gin.Context) {
		if !limiter.Allow() {
			c.AbortWithStatusJSON(429, gin.H{
				"status": 429,
				"error":  "Too Many Requests",
			})
			return
		}
		c.Next()
	}
}
