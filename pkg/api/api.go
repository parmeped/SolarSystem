package api

import (
	"github.com/gin-gonic/gin"
)

// Start a Gin server
func Start() *gin.Engine {

	router := gin.Default()
	v1 := router.Group("/v1")
	v1.GET("/dayClimate", getDayClimate())	

	return router
}

func getDayClimate() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}