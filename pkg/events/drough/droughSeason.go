package drought

// this package should expose methods to calculate the drought season.

import (
	"fmt"

	sol "github.com/SolarSystem/pkg/system"
	pos "github.com/SolarSystem/pkg/position"
)

type DroughtSeason struct {
	Name string
}

// TODO: Check func names and code. Code was moved, check if it works!

// GetDroughtSeasonsForYears returns the amount of droughts there's on a certain amount of years
// TODO: see if the hardcoded 365 value can be turned into a cfg call. What about leap years? This should receive a system, not positions
func GetDroughtSeasonsForYears(years int, sys *sol.System) int {
	return GetDroughtSeasonsForDays((years * 365), sys)
}

// GetDroughtSeasonsForDays returns the amount of droughts there's on a certain amount of days
func GetDroughtSeasonsForDays(days int, sys *sol.System) int {
	cycleDays := pos.TimeToSystemCycle(sys.Positions[0], sys.Positions[1], sys.Positions[2])
	multiplier := days / int(cycleDays)
	daysRemaining := days % int(cycleDays)

	// Register event & check
	droughEvent := DroughtSeason{"DroughtSeason"}	
	sys.NewEvent(droughEvent.Name)

	droughtSeasons, droughtDays := getDroughtSeasonsForCycle(int(cycleDays), sys.Positions)
	droughtSeasons = droughtSeasons * multiplier

	for _, v := range droughtDays {
		if v <= daysRemaining {
			droughtSeasons++
		} else {
			break
		}
	}
	sys.Events["OptimalAlignment"].AmountDays = droughtSeasons
	sys.Events["OptimalAlignment"].DaysEvent = append(sys.Events["OptimalAlignment"].DaysEvent, dayChecked)
	return droughtSeasons
}

// TODO: Maybe this could go inside the call to the other function, as a strategy or something.
// GetDroughtSeasonsForCycle calculates how many times there's a Drought season on a cycle.
func getDroughtSeasonsForCycle(cycleDays int, positions []*pos.Position) (int, []int) {
	// I know there's a least amount of time a couple of points can intersect. The check has to be for each pair of points
	fastestCycle, index, secondIndex := cycleDays, 0, 0

	// first get the two planets which are fastest to complete cycle. This way, it's the least we can check, for being the most certain.
	for k, v := range positions {
		if int(v.Planet.TimeToCycle) < fastestCycle {
			index = k
			fastestCycle = int(v.Planet.TimeToCycle)
		}
	}
	fastestCycle = cycleDays

	for k, v := range positions {
		if k != index && int(v.Planet.TimeToCycle) < fastestCycle {
			secondIndex = k
			fastestCycle = int(v.Planet.TimeToCycle)
		}
	}

	// Get where they meet at the start, then when they meet at any point, and how many times they meet. 
	timeStart, timeAny, amount := pos.GetTwoPointsIntersections(positions[index], positions[secondIndex])
	var intersect = pos.Intersections{timeStart, timeAny, amount, positions[index], positions[secondIndex]} // unkeyed fields?
	return checkForDroughts(intersect, cycleDays, positions)
}

// checks if there are droughts on a cycle. {amountOfDroughts, []daysOfDroughts}
func checkForDroughts(intersect pos.Intersections, cycleDays int, positions []*pos.Position) (int, []int) {

	// the period starts on a drought, since all planets start on pos 0
	amountOfDroughts, days := 1, int(intersect.TimeToFirst)
	positionToCheck, positionToCompare := &pos.Position{}, &pos.Position{}
	daysOfDroughts := []int{0} // TODO: this feels like a hack!

	// get the fastest planet
	if intersect.PositionA.Planet.TimeToCycle > intersect.PositionB.Planet.TimeToCycle {
		positionToCompare = intersect.PositionB
	} else {
		positionToCompare = intersect.PositionA
	}

	// get the position to check
	for _, v := range positions {
		if v != intersect.PositionA && v != intersect.PositionB {
			positionToCheck = v
		}
	}

	var positionPlanetToCheck, positionPlanetToCompare int
	for i := days; i < cycleDays; {
		positionPlanetToCheck = pos.GetPositionAtTime(&positionToCheck.Planet, i)
		positionPlanetToCompare = pos.GetPositionAtTime(&positionToCompare.Planet, i)
		i = i + days
		if checkPositionsForDrought(positionPlanetToCheck, positionPlanetToCompare) {
			daysOfDroughts = append(daysOfDroughts, i)
			amountOfDroughts++
		}
	}
	return amountOfDroughts, daysOfDroughts
}

// TODO: [Improvement] This could clearly use the collinearity function, but on the other hand this is a super small and simple checking method. 
// drought check helper. Compares two positions to find a drought. 
func checkPositionsForDrought(positionToCheck, positionToCompare int) bool {
	result := positionToCheck - positionToCompare
	fmt.Printf("posCheck: %v, posCompare: %v, result: %v \n", positionToCheck, positionToCompare, result)

	if result < 0 {
		result = result * -1
	}

	if result == 180 || result == 0 {
		return true
	} else {
		return false
	}
}

// TODO: check if this func is usefull
// Seems to be working fine.
func CheckForDrought(p1 *pos.Position, p2 *pos.Position) bool {
	angle, _ := pos.AngleBetweenPositions(p1, p2)
	//fmt.Printf("Drought Check Angle: %v \n", angle)

	if angle == 180 || p1.ClockWisePosition == p2.ClockWisePosition {
		return true
	} else {
		return false
	}
}
