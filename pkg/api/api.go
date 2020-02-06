package api

import (
	sol "github.com/SolarSystem/pkg/system"
	"github.com/gin-gonic/gin"
)

// IService is later implemented by the service
type IService interface {
	GetClimateForDay(day int) *sol.Day
}

// SetupRouter returns a Gin server
func SetupRouter(service IService) *gin.Engine {

	router := gin.Default()
	v1 := router.Group("/v1")

	v1.GET("/climateForDay/:Day", getClimateForDay(service))
	v1.GET("/", handle())

	return router
}
