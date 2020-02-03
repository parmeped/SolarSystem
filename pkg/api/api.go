package api

import (
	"github.com/gin-gonic/gin"
	sol "github.com/SolarSystem/pkg/system"
)

// IService is later implemented by the service
type IService interface {
	GetClimateForDay(day int) *sol.Day
}

// SetupRouter a Gin server
func SetupRouter(service IService) *gin.Engine {

	router := gin.Default()
	v1 := router.Group("/v1")
	v1.GET("/dayClimate", getClimateForDay(service))	

	return router
}
