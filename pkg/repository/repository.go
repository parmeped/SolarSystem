package repository

import (
	sol "github.com/SolarSystem/pkg/system"
	p "github.com/SolarSystem/pkg/planets"	
)

// Database wich contains the SolarSystem
type Database struct {
	SolarSystem *sol.System
}

// New Returns a pointer to a MockDb
func New() *Database {	
	sys := sol.New(p.GetPlanets())
	return &Database{sys}
}