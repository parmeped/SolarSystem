package events

import (
	"fmt"
	sol "github.com/SolarSystem/pkg/system"
	pos "github.com/SolarSystem/pkg/position"
	d "github.com/SolarSystem/pkg/events/drought"	
	o "github.com/SolarSystem/pkg/events/optimalalignment"	
	r "github.com/SolarSystem/pkg/events/rainseason"	
)

// TODO: This has to go..?
// func getDroughtSeasonsForYears(years int, sys *sol.System) int {
// 	return d.GetDroughtSeasonsForYears(years, sys)
// }

// func getDroughtSeasonsForDays(days int, sys *sol.System) int {
// 	return d.GetDroughtSeasonsForDays(days, sys)
// }

// func getOptimalAlignmentsForYears(years int, sys *sol.System) int {
// 	return o.GetOptimalAlignmentsForYears(years, sys)	
// }

// func getOptimalAlignmentsForDays(days int, sys *sol.System) int {
// 	return o.GetOptimalAlignmentsForDays(days, sys)
// }

// func getRainSeasonsForYears(years int, sys *sol.System) int {
// 	return r.GetRainSeasonsForDays(years, sys)
// }

// func getRainSeasonsForDays(days int, sys *sol.System) int {
// 	return r.GetRainSeasonsForDays(days, sys)
// }

// GetAmountPerEventForYears returns the amount of times an event happens on {years}
func GetAmountPerEventForYears(years int, sys *sol.System, event string) int {
	return GetAmountPerEventForDays(years * sys.Cfg.DaysInYear, sys, event)
}


// GetAmountPerEventForDays returns the amount of times an event happens on {days}. If event does not exist, returns -1
func GetAmountPerEventForDays(days int, sys *sol.System, event string) int {
	cycleDays := pos.TimeToSystemCycle(sys.Positions[0], sys.Positions[1], sys.Positions[2])
	multiplier := days / int(cycleDays)
	daysRemaining := days % int(cycleDays)

	if exists := registerEvent(event, sys); exists {
		// Reset as to renew event if conditions have changed
		sys.Events[event].AmountDays, sys.Events[event].DaysEvent = 0, []int{}

		eventAmount, eventDays := sys.Events[event].Implementations.GetEventPerCycle(int(cycleDays), sys)
		eventAmount = eventAmount * multiplier
	
		for _, v := range eventDays {
			if v <= daysRemaining {
				eventAmount++
			} else {
				break
			}
		}
		return eventAmount
	}
	return -1
}

// TODO: when event doesn't exist, this explodes.
// Registers the event in case it isn't already.
func registerEvent(event string, sys *sol.System) bool {
	if _, ok := sys.Events[event]; !ok {
		switch event {
		case "DroughtSeason": 
			d.RegisterEvent(sys)		
		case "OptimalAlignment": 
			o.RegisterEvent(sys)			
		case "RainSeason": 
			r.RegisterEvent(sys)
		default:
			fmt.Print("Event is not registered \n")			
			return false
		}		
	}
	return true
}

// GetConditionForDay gets the condition for a given day. 
func GetConditionForDay(day int, sys *sol.System) {

}