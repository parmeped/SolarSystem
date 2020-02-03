package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// TODO: have some validation!
func getClimateForDay(service IService) gin.HandlerFunc {
	return func(c *gin.Context) {

		request := c.Param("Day")

		if request == "" {
			c.JSON(http.StatusBadRequest, "Error, no parameter was provided")
			return
		}

		if day, err := strconv.ParseInt(request, 0, 0); err != nil {
			c.JSON(http.StatusInternalServerError, "There was an error processing the request")
			return
		} else {
			if day < 0 {
				c.JSON(http.StatusBadRequest, "Error, day queried must be equal to or greater than 0")
				return
			}
			response := service.GetClimateForDay(int(day))
			if response != nil {
				c.JSON(http.StatusOK, response)
			} else {
				c.JSON(http.StatusInternalServerError, nil)
			}
		}
	}
}
