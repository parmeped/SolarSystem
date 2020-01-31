package optimalalignment

// this package should expose methods to calculate the optimal alignment.

import (
	m "math"

	e "github.com/SolarSystem/pkg/events"
	pos "github.com/SolarSystem/pkg/position"
	sol "github.com/SolarSystem/pkg/system"
)

type optimalAlignment struct {
}

// GetOptimalAlignmentsForYears returns how many optimal climate alignments happen in {n} years
func GetOptimalAlignmentsForYears(years int, sys *sol.System) int {
	return GetOptimalAlignmentsForDays((years * 365), sys)
}

// TODO: This logic is the same as the other events. This is a candidate for a generic, or at least a strategy
// GetOptimalAlignmentsForDays returns how many optimal climate alignments happen in {n} days
func GetOptimalAlignmentsForDays(days int, sys *sol.System) int {
	cycleDays := pos.TimeToSystemCycle(sys.Positions[0], sys.Positions[1], sys.Positions[2])
	multiplier := days / int(cycleDays)
	daysRemaining := days % int(cycleDays)

	optimalSeasons, optimalDays := getOptimalAlignmentsForCycle(int(cycleDays), sys)
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
func getOptimalAlignmentsForCycle(cycleDays int, sys *sol.System) (int, []int) {
	// for each day on the cycle, see if there's an alignment. how can you pass a function that's executed after each day it rotates?

	optAlignment := optimalAlignment{}
	event := e.New("OptimalAlignment")
	sys.AddEvent(event)

	passFunction := []sol.IExecute{} // small compromise so that RotateAndExecute can receive more than 1 function at a time. In theory
	passFunction = append(passFunction, &optAlignment)

	sol.RotateAndExecute(cycleDays, sys, passFunction)

	return 0, []int{}
}

// TODO: how can I retrieve this event?
// Middleware function for Polymorfism implementation
func DailyCheck(sys *sol.System) {
	days := &(*sys.Events)["OptimalAlignment"]
}

// Checks if {n} positions are aligned.
func checkAlignmentForPositions(positions []*pos.Position) bool {
	coordinates := []pos.Coordinate{}
	for _, v := range positions {
		coordinates = append(coordinates, pos.ConvertPolarToCartesian(v))
	}
	return checkAlignmentForCoordinates(&coordinates)
}

// TODO: maybe when returning true, should also check for alignment with sun
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
