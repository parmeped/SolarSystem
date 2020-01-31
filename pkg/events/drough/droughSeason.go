package drough

// this package should expose methods to calculate the drough season.

import (
	"fmt"

	pos "github.com/SolarSystem/pkg/position"
)

// TODO: Check func names and code. Code was moved, check if it works!

// GetDroughSeasonsForYears returns the amount of droughs there's on a certain amount of years
// TODO: see if the hardcoded 365 value can be turned into a cfg call. What about leap years? This should receive a system, not positions
func GetDroughSeasonsForYears(years int, positions []*pos.Position) (int, []int) {
	return GetDroughSeasonsForDays((years * 365), positions)
}

// GetDroughSeasonsForDays returns the amount of droughs there's on a certain amount of days
func GetDroughSeasonsForDays(days int, positions []*pos.Position) (int, []int) {
	cycleDays := pos.TimeToSystemCycle(positions[0], positions[1], positions[2])
	multiplier := days / int(cycleDays)
	daysRemaining := days % int(cycleDays)
	droughSeasons, droughDays := getDroughSeasonsForCycle(int(cycleDays), positions)
	droughSeasons = droughSeasons * multiplier

	for _, v := range droughDays {
		if v <= daysRemaining {
			droughSeasons++
		} else {
			break
		}
	}
	return droughSeasons, droughDays
}

// TODO: Maybe this could go inside the call to the other function, as a strategy or something.
// GetDroughSeasonsForCycle calculates how many times there's a Drough season on a cycle.
func getDroughSeasonsForCycle(cycleDays int, positions []*pos.Position) (int, []int) {
	// I know there's a least amount of time a couple of points can intersect. The check has to be for each pair of points
	fastestCycle, index, secondIndex := cycleDays, 0, 0

	// first get the two planets which are fastest to complete cycle
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

	timeStart, timeAny, amount := pos.GetTwoPointsIntersections(positions[index], positions[secondIndex])
	var intersect = pos.Intersections{timeStart, timeAny, amount, positions[index], positions[secondIndex]} // unkeyed fields?
	return checkForDroughs(intersect, cycleDays, positions)
}

// checks if there are droughs on a cycle. {amountOfDroughs, []daysOfDroughs}
func checkForDroughs(intersect pos.Intersections, cycleDays int, positions []*pos.Position) (int, []int) {

	// the period starts on a drough, since all planets start on pos 0
	amountOfDroughs, days := 1, int(intersect.TimeToFirst)
	positionToCheck, positionToCompare := &pos.Position{}, &pos.Position{}
	daysOfDroughs := []int{0} // TODO: this feels like a hack!

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
		if checkPositionsForDrough(positionPlanetToCheck, positionPlanetToCompare) {
			daysOfDroughs = append(daysOfDroughs, i)
			amountOfDroughs++
		}
	}
	return amountOfDroughs, daysOfDroughs
}

// drough check helper. Compares two positions to find a drough
func checkPositionsForDrough(positionToCheck, positionToCompare int) bool {
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
func CheckForDrough(p1 *pos.Position, p2 *pos.Position) bool {
	angle, _ := pos.AngleBetweenPositions(p1, p2)
	//fmt.Printf("Drough Check Angle: %v \n", angle)

	if angle == 180 || p1.ClockWisePosition == p2.ClockWisePosition {
		return true
	} else {
		return false
	}
}
