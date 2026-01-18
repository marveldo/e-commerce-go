package routes

import (
	"github.com/gin-gonic/gin"
	)

func GetEngine() *gin.Engine {
	g := gin.Default()

	return g
}