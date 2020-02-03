package rainseason

import (
	m "math"
	
	pos "github.com/SolarSystem/pkg/position"
	sol "github.com/SolarSystem/pkg/system"
)

// RainSeason base
type RainSeason struct {
	Name                    string	
	ShouldCheckForAlignment bool
	NextDayToAvoid          int
}

// Registers the event on the system.
func RegisterEvent(sys *sol.System) {
	// Register event & check.
	rainSeasonEvent := RainSeason{"RainSeason", true, 0}
	//sys.AddCheck(rainSeasonEvent)
	sys.NewEvent(rainSeasonEvent.Name, rainSeasonEvent, true)
	// If we've already checked for droughtSeasons, can avoid checking for alignment.
	// this check has to be done because the formula used gives true when all 4 points are aligned, which therefore isn't a triangle.
	if val, ok := sys.Events["DroughtSeason"]; ok {
		if len(val.DaysEvent) > 0 {
			rainSeasonEvent.ShouldCheckForAlignment = false
		}
	}
} 

// GetEventPerCycle checks if there are Rainseasons on a cycle. Implementation of IEvent
func (event RainSeason) GetEventPerCycle(cycleDays int, sys *sol.System) (int, []int) {
	
	sol.RotateAndExecute(cycleDays, sys, event.Name)
	return sys.Events["RainSeason"].AmountDays, sys.Events["RainSeason"].DaysEvent
}

// DailyCheck function used for daily checks on a system after it rotates one day
func (event RainSeason) DailyCheck(sys *sol.System, dayChecked int) {
	avoidCheck := false
	if event.ShouldCheckForAlignment {
		
		// this check has to be done because the formula used gives true when all 4 points are aligned, which therefore isn't a triangle.
		coords := pos.ConvertPolarSliceToCartesian(sys.Positions)
		*coords = append(*coords, *sys.SunCoordinates)
		avoidCheck, _ = pos.CheckAlignmentForCoordinates(coords)
	} else {

		// If drought season event has already been calculated, the days can be reused.
		dayToAvoid := event.NextDayToAvoid
		if dayChecked == sys.Events["DroughtSeason"].DaysEvent[dayToAvoid] {
			avoidCheck = true

			if dayToAvoid < len(sys.Events["DroughSeason"].DaysEvent)-1 {
				event.NextDayToAvoid = sys.Events["DroughSeason"].DaysEvent[dayToAvoid+1]
			}			
		}
	}

	if !avoidCheck {
		if isInside, coords := pos.ConvertToCartesianAndExecute(sys.Positions, checkSunInsideTriangle); isInside {
			sys.Events["RainSeason"].AmountDays++
			sys.Events["RainSeason"].DaysEvent = append(sys.Events["RainSeason"].DaysEvent, dayChecked)
			
			if peakRainDay, newPerimeter := checkForPeakRainDay(coords, sys.Events["RainSeason"].MaxPerimeter); peakRainDay {
				sys.Events["RainSeason"].PeakDay = dayChecked						
				sys.Events["RainSeason"].MaxPerimeter = newPerimeter
			}
		}
	}
}

func checkForPeakRainDay(coords *[]pos.Coordinate, perimeter float32) (bool, float32) {
	c1, c2, c3 := (*coords)[0], (*coords)[1], (*coords)[2]
	var sideA, sideB, sideC float64
	sideA = m.Sqrt(m.Pow(float64(c2.X-c3.X), 2) + m.Pow(float64(c2.Y-c3.Y), 2))
	sideB = m.Sqrt(m.Pow(float64(c1.X-c3.X), 2) + m.Pow(float64(c1.Y-c3.Y), 2))
	sideC = m.Sqrt(m.Pow(float64(c1.X-c2.X), 2) + m.Pow(float64(c1.Y-c2.Y), 2))
	return (sideA + sideB + sideC) > float64(perimeter), float32(sideA + sideB + sideC)
}

func checkSunInsideTriangle(coords *[]pos.Coordinate) (bool, *[]pos.Coordinate) {
	// sun coordinates
	sunCoords := pos.Coordinate{nil, 0, 0}

	b1 := sign(sunCoords, (*coords)[0], (*coords)[1]) < 0
	b2 := sign(sunCoords, (*coords)[1], (*coords)[2]) < 0
	b3 := sign(sunCoords, (*coords)[2], (*coords)[0]) < 0
	return (b1 == b2) && (b2 == b3), coords
}

// helper func 
func sign(c1, c2, c3 pos.Coordinate) float32 {
	return (c1.X-c3.X)*(c2.Y-c3.Y) - (c2.X-c3.X)*(c1.Y-c3.Y)
}
