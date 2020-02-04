package config

import (

	"time"
	"fmt"
	er "github.com/SolarSystem/pkg/utl/error"
	pRepo "github.com/SolarSystem/pkg/planets"	
)

// I believe daysInYear should always be higher. First, since a year is defined as the time it takes 
const (
	daysInYear 	 = 365  	
	orbit 		 = 360
)

// Load returns Configuration struct
func Load() *Configuration {
	er.HandleError("Load Config")
	
	timeStamp()

	return &Configuration{
		pRepo.GetPlanets(orbit),
		daysInYear,
		time.Now(),
		orbit,
	}		
}

// GetOrbit gives the orbit
// TODO: [Improvement] This shouldn't be loaded from here, but from the config after load.
func GetOrbit() int {
	return orbit
}

// Configuration holds data necessery for configuring application
type Configuration struct {		
	Planets	         []pRepo.Planet
	DaysInYear       int    	  
	StartingDate     time.Time    
	Orbit			 int	
}


func timeStamp() {
	fmt.Print("////////////////   Program starting... ////////////////// \n")
	fmt.Printf("Time: %v \n", time.Now())
	fmt.Print("-------------------------------\n")

}