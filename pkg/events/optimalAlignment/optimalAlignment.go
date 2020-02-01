package optimalalignment

// this package should expose methods to calculate the optimal alignment.

import (
	"math"
	"fmt"

	pos "github.com/SolarSystem/pkg/position"
	sol "github.com/SolarSystem/pkg/system"
)

type optimalAlignment struct {
	Name string
}

// TODO: here, as in other methods, we should make a days calculator as to get amount of days on years. 
// GetOptimalAlignmentsForYears returns how many optimal climate alignments happen in {n} years
func GetOptimalAlignmentsForYears(years int, sys *sol.System) int {
	return GetOptimalAlignmentsForDays((years * 365), sys)
}

// TODO: This logic is the same as the other events. This is a candidate for a generic, or at least a strategy.
// GetOptimalAlignmentsForDays returns how many optimal climate alignments happen in {n} days
func GetOptimalAlignmentsForDays(days int, sys *sol.System) int {
	cycleDays := pos.TimeToSystemCycle(sys.Positions[0], sys.Positions[1], sys.Positions[2])
	multiplier := days / int(cycleDays)
	daysRemaining := days % int(cycleDays)

	optimalSeasons, optimalDays := getOptimalAlignmentsForCycle(int(cycleDays), sys)
	optimalSeasons = optimalSeasons * multiplier

	// see if there's more events happening 
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
	
	// Register event & check
	optAlignment := optimalAlignment{"OptimalAlignment"}
	sys.AddCheck(optAlignment)
	sys.NewEvent(optAlignment.Name)

	// Execute a complete cycle check on the system. 
	sol.RotateAndExecute(cycleDays, sys)

	// The amount and the days should be already loaded
	return sys.Events["OptimalAlignment"].AmountDays, sys.Events["OptimalAlignment"].DaysEvent
}

// daily check function used for daily checks on a system after it rotates one day
func (opt optimalAlignment) DailyCheck(sys *sol.System, dayChecked int) {
	isAligned, coords := checkAlignmentForPositions(sys.Positions)
	if isAligned {
		// sun coordinates to check if planets are also aligned with the sun. 		
		*coords = append(*coords, *sys.SunCoordinates)
		if isAligned, _ = checkAlignmentForCoordinates(coords); !isAligned {
			fmt.Printf("Day: %v, Coordinates: %v \n", dayChecked, coords)
			sys.Events["OptimalAlignment"].AmountDays++
			sys.Events["OptimalAlignment"].DaysEvent = append(sys.Events["OptimalAlignment"].DaysEvent, dayChecked)
		}

	}
}

// Checks if {n} positions are aligned. Also returns the positions converted to coordinates
func checkAlignmentForPositions(positions []*pos.Position) (bool, *[]pos.Coordinate) {
	coordinates := []pos.Coordinate{}
	for _, v := range positions {
		coordinates = append(coordinates, pos.ConvertPolarToCartesian(v))
	}
	return checkAlignmentForCoordinates(&coordinates)
}


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
		coord3 := (*coords)[k+2] 

		// find slopes of coordinates and check between a common point
		slopeA := float64((coord1.Y - coord3.Y) / (coord1.X - coord3.X))
		slopeB := float64((coord3.Y - coord2.Y) / (coord3.X - coord2.X))


		// TODO: [Improvement] maybe improve time frames so that this check is more precise.
		// this rounds to later determine collinearity. Since we check once per day, this is the compromise being made. 
		// It would work better if checks were done on smaller time frames, but would have a higher impact on performance. 
		slopeA, slopeB = math.Floor(slopeA*10)/10, math.Floor(slopeB*10)/10

		if slopeA == slopeB {

			aligned = true
			// would break otherwise, since all points have been checked
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
