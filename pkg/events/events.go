package events

import (
	"fmt"
	"time"

	d "github.com/SolarSystem/pkg/events/drought"
	o "github.com/SolarSystem/pkg/events/optimalalignment"
	r "github.com/SolarSystem/pkg/events/rainseason"
	pos "github.com/SolarSystem/pkg/position"
	sol "github.com/SolarSystem/pkg/system"
	er "github.com/SolarSystem/pkg/utl/error"
)

// GetAmountPerEventForYears returns the amount of times an event happens on {years}
func GetAmountPerEventForYears(years int, sys *sol.System, event string) int {
	defer er.HandleError("GetAmountPerEventForYears")

	return GetAmountPerEventForDays(years*sys.Cfg.DaysInYear, sys, event)
}

// GetAmountPerEventForDays returns the amount of times an event happens on {days}. If event does not exist, returns -1
func GetAmountPerEventForDays(days int, sys *sol.System, event string) int {
	defer er.HandleError("GetAmountPerEventForDays")

	cycleDays := pos.TimeToSystemCycle(sys.Positions[0], sys.Positions[1], sys.Positions[2])
	multiplier := days / int(cycleDays)
	daysRemaining := days % int(cycleDays)
	sys.TimeToCycle = cycleDays

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

// Registers the event in case it isn't already.
func registerEvent(event string, sys *sol.System) bool {
	defer er.HandleError("registerEvent")


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
func GetConditionForDay(sys *sol.System, day int) *sol.Day {
	defer er.HandleError("GetConditionForDay")

	if day < 0 {
		return nil
	}

	date := &sol.Day{day, sys.Cfg.StartingDate.AddDate(0, 0, day).Format(time.RFC850), sol.Normal}

	// No other way to get the peak rain day. Should check anyway for each point.
	GetAmountPerEventForYears(1, sys, "RainSeason")
	rainy, peakRain := r.SingleDayCheck(sys, day)
	if rainy {
		if peakRain {
			date.Condition = sol.PeakRain
		} else {
			date.Condition = sol.Rainy
		}
		return date
	}
	if drought := d.SingleDayCheck(sys, day); drought {
		date.Condition = sol.Drought
		return date
	}
	if optimal := o.SingleDayCheck(sys, day); optimal {
		date.Condition = sol.Optimal
		return date
	}
	return date
}
