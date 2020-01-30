package events

// this package should expose methods to calculate the optimal alignment.

import (
	m "math"

	pos "github.com/SolarSystem/pkg/position"
	sol "github.com/SolarSystem/pkg/system"
)

// GetOptimalAlignmentsForYears returns how many optimal climate alignments happen in {n} years
func GetOptimalAlignmentsForYears(years int, sys *sol.System) int {
	return GetOptimalAlignmentsForDays((years * 365), sys)
}

// TODO: This logic is the same as the other events. This is a candidate for a generic, or at least a strategy
// GetOptimalAlignmentsForDays returns how many optimal climate alignments happen in {n} days
func GetOptimalAlignmentsForDays(days int, sys *sol.System) int {
	positions := sys.Positions
	cycleDays := pos.TimeToSystemCycle(positions[0], positions[1], positions[2])
	multiplier := days / int(cycleDays)
	daysRemaining := days % int(cycleDays)

	optimalSeasons, optimalDays := getOptimalAlignmentsForCycle(int(cycleDays), positions)
	optimalSeasons = optimalSeasons * multiplier

	for _, v := range optimalDays {
		if v <= daysRemaining {
			optimalSeasons++
		} else {
			break
		}
	}
	return optimalSeasons
}

// checks if there are optimalAlignments on a cycle. {amountOfOptimals, []daysOfOptimals}
func getOptimalAlignmentsForCycle(cycleDays int, positions []*pos.Position) (int, []int) {
	return 0, []int{}
}

// Checks if {n} positions are aligned.
func CheckAlignmentForPositions(positions []*pos.Position) bool {
	coordinates := []pos.Coordinate{}
	for _, v := range positions {
		coordinates = append(coordinates, pos.ConvertPolarToCartesian(v))
	}
	return checkAlignmentForCoordinates(&coordinates)
}

// TODO: Somewhere here we have to NOT check on drough days
// checks if {n} coordinates are aligned.
func checkAlignmentForCoordinates(coords *[]pos.Coordinate) bool {
	aligned := false
	coordsLength := len(*coords)

	// simple validation check
	if coordsLength < 3 {
		return false
	}

	// this should cycle through all of them
	for k, _ := range *coords {
		coord1 := (*coords)[k]
		coord2 := (*coords)[k+1]
		coord3 := (*coords)[k+2] // this is the common point checked.

		// find slopes of coordinates and check between a common point
		slopeA := float64((coord1.Y - coord3.Y) / (coord1.X - coord3.X))
		slopeB := float64((coord3.Y - coord2.Y) / (coord3.X - coord2.X))

		// could convert to int for slope checking and performance gain?
		if m.Round(slopeA) == m.Round(slopeB) {

			aligned = true
			// would break otherwise, and all points have been checked
			if k+3 == coordsLength {
				break
			}
		} else {
			aligned = false
			break
		}
	}
	return aligned
}
