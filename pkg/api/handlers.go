package api

import (
	"net/http"	
	"github.com/gin-gonic/gin"
)


// TODO: have some validation!
func getClimateForDay(service IService) gin.HandlerFunc {
	return func(c *gin.Context) {
		day := service.GetClimateForDay(0)
		if day != nil {
			c.JSON(http.StatusOK, day)
		} else {
			c.JSON(http.StatusInternalServerError, nil)
		}
	}
}