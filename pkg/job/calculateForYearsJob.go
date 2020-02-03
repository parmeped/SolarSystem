package job

import (
		
	"time"
	sol "github.com/SolarSystem/pkg/system"
	e "github.com/SolarSystem/pkg/events"
	repo "github.com/SolarSystem/pkg/repository"
)

// CalculateModelForYears fills the repository with the conditions the planet will have on each day up to a certain date, given an amount of {years}
func CalculateModelForYears(years int, sys *sol.System, db *repo.Database) {
	// Run the model for every event, 1 year.
	e.GetAmountPerEventForYears(1, sys, "DroughtSeason")
	e.GetAmountPerEventForYears(1, sys, "OptimalAlignment")
	e.GetAmountPerEventForYears(1, sys, "RainSeason")
	i, startingDate, startingDay := 0, sys.Cfg.StartingDate, 0
	for i < years {
		calculateModelForAYear(startingDate,  sys, startingDay, db)
		startingDate = startingDate.AddDate(i+1,0,0)
		startingDay += sys.Cfg.DaysInYear
		i++
	}
}

// Calculates the model for a year and adds it to the repository
func calculateModelForAYear(date time.Time, sys *sol.System, dayID int, db *repo.Database) {
	daysInYear, days := sys.Cfg.DaysInYear, []*sol.Day{}
	e := sys.Events
	condition := ""
	droughtDays, rainDays, optimalDays, peakRainDay := 
	e["DroughtSeason"].DaysEvent, e["RainSeason"].DaysEvent, e["OptimalAlignment"].DaysEvent, e["RainSeason"].PeakDay

	for i := 0; i < daysInYear; i++ {
		droughtDays, rainDays, optimalDays, condition = checkCondition(i, peakRainDay, droughtDays, rainDays, optimalDays)
		day := &sol.Day{dayID, date.Format(time.RFC850), condition}
		days = append(days, day)
		dayID++
		date = date.AddDate(0,0,1)
	}
	db.AddDaysModel(days)
}


// check for every day, crop slice when one is found as to improve performance on later checks.
func checkCondition(dayNumber, peakDay int, days ...[]int) ([]int, []int, []int, string) {
	condition := ""
	if dayNumber == peakDay {
		return days[0], days[1], days[2], sol.PeakRain
	}
	// First one is drough
	for k, v := range days[0] {
		if v == dayNumber {
			condition = sol.Drought			
			return days[0][k+1:], days[1], days[2], condition
		}
	}
	// Second is Rain
	for k, v := range days[1] {
		if v == dayNumber {
			condition = sol.Rainy			
			return days[0], days[1][k+1:], days[2], condition
		}
	}
	// Third one is Optimal days
	for k, v := range days[2] {
		if v == dayNumber {
			condition = sol.Optimal			
			return days[0], days[1], days[2][k+1:], condition
		}
	}
	return 	days[0], days[1], days[2], sol.Normal
}