package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	processingError      = "There was an error processing the request"
	noParameterError     = "Error, no parameter was provided"
	dayGreaterThan0Error = "Error, day queried must be equal to or greater than 0"
	nothingToSee         = "Nothing to see here! Move along..."
)

// Default
func handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, nothingToSee)
	}
}

// GetClimateForDay handler
func getClimateForDay(service IService) gin.HandlerFunc {
	return func(c *gin.Context) {

		request := c.Param("Day")

		if request == "" {
			c.JSON(http.StatusBadRequest, noParameterError)
			return
		}

		if day, err := strconv.ParseInt(request, 0, 0); err != nil {
			c.JSON(http.StatusInternalServerError, processingError)
			return
		} else {
			if day < 0 {
				c.JSON(http.StatusBadRequest, dayGreaterThan0Error)
				return
			}
			response := service.GetClimateForDay(int(day))
			if response != nil {
				c.JSON(http.StatusOK, response)
			} else {
				c.JSON(http.StatusInternalServerError, processingError)
			}
		}
	}
}
