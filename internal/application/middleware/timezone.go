package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
)

func Gettimezone() gin.HandlerFunc {
	return func(c *gin.Context) {
	  var tz *time.Location
	  tzHeader := c.GetHeader("X-Timezone")
	  if tzHeader != "" {
		tz , _ = time.LoadLocation(tzHeader)
	  }
	  if tz == nil {
		tz = time.UTC
	  }
	  c.Set("timezone", tz)
	  c.Next()

	}
}