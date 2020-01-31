package optimalalignment

// this package should expose methods to calculate the optimal alignment.

import (
	pos "github.com/SolarSystem/pkg/position"
	sol "github.com/SolarSystem/pkg/system"
)

type optimalAlignment struct {
	Name string
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

	optimalSeasons, optimalDays := GetOptimalAlignmentsForCycle(int(cycleDays), sys)
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
func GetOptimalAlignmentsForCycle(cycleDays int, sys *sol.System) (int, []int) {
	// for each day on the cycle, see if there's an alignment. how can you pass a function that's executed after each day it rotates?

	optAlignment := optimalAlignment{"OptimalAlignment"}
	sys.AddCheck(optAlignment)
	sys.NewEvent(optAlignment.Name)

	sol.RotateAndExecute(cycleDays, sys)

	return 0, []int{}
}

// daily check function used for daily checks on a system after it rotates one day
func (opt optimalAlignment) DailyCheck(sys *sol.System, dayChecked int) {
	isAligned, coords := CheckAlignmentForPositions(sys.Positions)
	if isAligned {
		// sun coordinate
		sunCoord := pos.Coordinate{}
		sunCoord.X, sunCoord.Y = 0, 0
		*coords = append(*coords, sunCoord)
		if isAligned, _ = checkAlignmentForCoordinates(coords); !isAligned {
			sys.Events["OptimalAlignment"].AmountDays++
			sys.Events["OptimalAlignment"].DaysEvent = append(sys.Events["OptimalAlignment"].DaysEvent, dayChecked)
		}

	}
}

// Checks if {n} positions are aligned. Also returns the positions converted to coordinates
func CheckAlignmentForPositions(positions []*pos.Position) (bool, *[]pos.Coordinate) {
	coordinates := []pos.Coordinate{}
	for _, v := range positions {
		coordinates = append(coordinates, pos.ConvertPolarToCartesian(v))
	}
	return checkAlignmentForCoordinates(&coordinates)
}

// TODO: maybe when returning true, should also check for alignment with sun
// TODO: Somewhere here we have to NOT check on drough days
// checks if {n} coordinates are aligned.
func checkAlignmentForCoordinates(coords *[]pos.Coordinate) (bool, *[]pos.Coordinate) {
	aligned := false
	coordsLength := len(*coords)

	// simple validation check
	if coordsLength < 3 {
		return false, coords
	}

	// this should cycle through all of them
	for k := range *coords {
		coord1 := (*coords)[k]
		coord2 := (*coords)[k+1]
		coord3 := (*coords)[k+2] // this is the common point checked.

		// // find slopes of coordinates and check between a common point
		// slopeA := (coord1.Y - coord3.Y) / (coord1.X - coord3.X)
		// slopeB := (coord3.Y - coord2.Y) / (coord3.X - coord2.X)

		slope := coord1.X*(coord2.Y-coord3.Y) +
			coord2.X*(coord3.Y-coord1.Y) +
			coord3.X*(coord1.Y-coord2.Y)

		//slopeA, slopeB = m.Floor(slopeA*100)/100, m.Floor(slopeB*100)/100

		// could convert to int, for slope checking and performance gain?
		if slope == 0 {

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

	return aligned, coords
}
