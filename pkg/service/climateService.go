package service

import (
	"github.com/SolarSystem/pkg/api"
	repo "github.com/SolarSystem/pkg/repository"
	sol "github.com/SolarSystem/pkg/system"
)
	

// ClimateService base
type ClimateService struct {	
	DB	*repo.Database
}

// New returns a pointer to a climate service
func New(DB *repo.Database) api.IService {
	return ClimateService{DB}
}

// GetClimateForDay implementation for returning the climate of a certain day
func (service ClimateService) GetClimateForDay(day int) *sol.Day {
	if val, ok := service.DB.Days[day]; ok {
		return val
	} else {
		return nil	
	}
}