package validator

import (
	// "log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)


func Validate[T interface{}] ( g *gin.Context, v interface{}, d  *T) *T {
	
    err := g.ShouldBindBodyWithJSON(v)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"status" : http.StatusBadRequest,
			"error" : err.Error(),
		})
		return nil
	}
	err = copier.Copy(d , v)
	if err != nil {
		g.JSON(http.StatusInternalServerError , gin.H{
			"status" : http.StatusInternalServerError,
			"error" : err.Error(),
		})
		return nil
	}
	return d
}

