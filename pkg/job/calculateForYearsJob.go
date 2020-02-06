package job

import (
	"time"

	e "github.com/SolarSystem/pkg/events"
	repo "github.com/SolarSystem/pkg/repository"
	sol "github.com/SolarSystem/pkg/system"
	er "github.com/SolarSystem/pkg/utl/error"
)

// CalculateModelForYears fills the repository with the conditions the planet will have on each day up to a certain date, given an amount of {years}
func CalculateModelForYearsJob(years int, sys *sol.System, db *repo.Database) {
	er.HandleError("CalculateModelForYearsJob")

	// Run the model for every event, 1 year.
	e.GetAmountPerEventForYears(1, sys, "DroughtSeason")
	e.GetAmountPerEventForYears(1, sys, "OptimalAlignment")
	e.GetAmountPerEventForYears(1, sys, "RainSeason")
	i, startingDate, startingDay, remainingDays := 0, sys.Cfg.StartingDate, 0, years*sys.Cfg.DaysInYear%sys.TimeToCycle
	for i < years {
		calculateModelForDays(startingDate, sys, startingDay, sys.TimeToCycle, db)
		startingDate = startingDate.AddDate(0, 0, sys.TimeToCycle)
		startingDay += sys.TimeToCycle
		i++
	}
	if remainingDays > 0 {
		calculateModelForDays(startingDate, sys, startingDay, remainingDays, db)
	}
}

// Calculates the model for a year and adds it to the repository
func calculateModelForDays(date time.Time, sys *sol.System, dayID, limit int, db *repo.Database) {
	er.HandleError("calculateModelForDaysJob")

	days := []*sol.Day{}
	e := sys.Events
	condition := ""
	droughtDays, rainDays, optimalDays, peakRainDay :=
		e["DroughtSeason"].DaysEvent, e["RainSeason"].DaysEvent, e["OptimalAlignment"].DaysEvent, e["RainSeason"].PeakDay

	for i := 0; i < limit; i++ {
		droughtDays, rainDays, optimalDays, condition = checkCondition(i, peakRainDay, droughtDays, rainDays, optimalDays)
		day := &sol.Day{dayID, date.Format(time.RFC850), condition}
		days = append(days, day)
		dayID++
		date = date.AddDate(0, 0, 1)
	}
	db.AddDaysModel(days)
}

// check for every day, crop slice when one is found as to improve performance on later checks.
func checkCondition(dayNumber, peakDay int, days ...[]int) ([]int, []int, []int, string) {
	er.HandleError("checkConditionJob")

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
	return days[0], days[1], days[2], sol.Normal
}
