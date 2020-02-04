package drought

import (

	er "github.com/SolarSystem/pkg/utl/error"
	pos "github.com/SolarSystem/pkg/position"
	sol "github.com/SolarSystem/pkg/system"
)

// DroughtSeason base
type DroughtSeason struct {
	Name string
}

// RegisterEvent registers the event on the system.
func RegisterEvent(sys *sol.System) {
	// Register event.
	droughEvent := DroughtSeason{"DroughtSeason"}
	sys.NewEvent(droughEvent.Name, droughEvent, false)
}

// DailyCheck function used for daily checks on a system after it rotates one day
func (event DroughtSeason) DailyCheck(sys *sol.System, dayChecked int) {
	// Does not need to perform a daily check.
}

// SingleDayCheck to know if planets are aligned with sun.
func SingleDayCheck(sys *sol.System, dayChecked int) bool {
	er.HandleError("SingleDayCheckDR")

	for _, v := range sys.Positions {
		v.ClockWisePosition = pos.GetPositionAtTime(&v.Planet, dayChecked)
	}
	coords := pos.ConvertPolarSliceToCartesian(sys.Positions)
	*coords = append(*coords, *sys.SunCoordinates)
	alignedWithSun, _ := pos.CheckAlignmentForCoordinates(coords)
	return alignedWithSun
}

// GetEventPerCycle calculates how many times there's a Drought season on a cycle. Implementation of IEvent
func (event DroughtSeason) GetEventPerCycle(cycleDays int, sys *sol.System) (int, []int) {
	er.HandleError("GetEventPerCycleDR")

	fastestCycle, firstIndex, secondIndex, lastIndex := cycleDays, 0, 0, 0
	positions := sys.Positions

	// first get the two planets which are fastest to complete cycle. This way, it's the least we can check, for being the most certain.
	for k, v := range positions {
		if int(v.Planet.OrbitalPeriod) < fastestCycle {
			firstIndex = k
			fastestCycle = int(v.Planet.OrbitalPeriod)
		}
	}
	fastestCycle = cycleDays

	for k, v := range positions {
		if k != firstIndex && int(v.Planet.OrbitalPeriod) < fastestCycle {
			secondIndex = k
			fastestCycle = int(v.Planet.OrbitalPeriod)
		} else {
			// slowest planet
			lastIndex = k
		}
	}

	// Get where they meet at the start, then when they meet at any point, and how many times they meet.
	_, timeAny, _ := pos.GetTwoPointsIntersections(positions[firstIndex], positions[secondIndex])

	droughtSeasons, daysOfDroughts := checkForDroughts(int(timeAny), firstIndex, lastIndex, cycleDays, sys)

	sys.Events["DroughtSeason"].AmountDays = droughtSeasons
	sys.Events["DroughtSeason"].DaysEvent = daysOfDroughts
	return droughtSeasons, daysOfDroughts
}

// checks if there are droughts on a cycle. {amountOfDroughts, []daysOfDroughts}.
func checkForDroughts(daysToFirst, firstIndex, lastIndex, cycleDays int, sys *sol.System) (int, []int) {
	er.HandleError("checkForDroughts")

	amountOfDroughts, daysOfDroughts := 0, []int{}
	positionToCheck, positionToCompare := sys.Positions[firstIndex], sys.Positions[lastIndex]

	// days is the amount of days it takes for the fastests planets to meet again. Therefore can safely check each time they cross, for the position of the last one
	var positionPlanetToCheck, positionPlanetToCompare int
	for i := 0; i < cycleDays; {
		positionPlanetToCheck = pos.GetPositionAtTime(&positionToCheck.Planet, i)
		positionPlanetToCompare = pos.GetPositionAtTime(&positionToCompare.Planet, i)
		if ok := checkPositionsForDrought(positionPlanetToCheck, positionPlanetToCompare, sys.Cfg.Orbit); ok {
			daysOfDroughts = append(daysOfDroughts, i)
			amountOfDroughts++
		}
		i = i + daysToFirst
	}
	return amountOfDroughts, daysOfDroughts
}

// TODO: [Improvement] This could clearly use the collinearity function, but on the other hand this is a super small and simple checking method.
// drought check helper. Compares two positions to find a drought.
func checkPositionsForDrought(positionToCheck, positionToCompare, orbitalLength int) bool {
	er.HandleError("checkPositionsForDrought")

	result := float32(positionToCheck) - float32(positionToCompare)

	if result < 0 {
		result = result * -1
	}

	if result == float32(orbitalLength/2) || result == 0 {
		return true
	}
	return false

}
