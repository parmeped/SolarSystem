package optimalalignment

// this package should expose methods to calculate the optimal alignment.

import (
	pos "github.com/SolarSystem/pkg/position"
	sol "github.com/SolarSystem/pkg/system"
)

// OptimalAlignment base
type OptimalAlignment struct {
	Name string
}

// RegisterEvent registers the event on the system.
func RegisterEvent(sys *sol.System) {
	// Register event & check
	optAlignEvent := OptimalAlignment{"OptimalAlignment"}	
	//sys.AddCheck(optAlignEvent)
	sys.NewEvent(optAlignEvent.Name, optAlignEvent, true)
} 

// GetEventPerCycle checks if there are optimalAlignments on a cycle. Implementation of IEvent
func (event OptimalAlignment) GetEventPerCycle(cycleDays int, sys *sol.System) (int, []int) {	
	
	sol.RotateAndExecute(cycleDays, sys, event.Name)
	return sys.Events["OptimalAlignment"].AmountDays, sys.Events["OptimalAlignment"].DaysEvent
}

// DailyCheck function used for daily checks on a system after it rotates one day
func (event OptimalAlignment) DailyCheck(sys *sol.System, dayChecked int) {
	isAligned, coords := pos.ConvertToCartesianAndExecute(sys.Positions, pos.CheckAlignmentForCoordinates)
	if isAligned {		
		*coords = append(*coords, *sys.SunCoordinates)
		if isAligned, _ = pos.CheckAlignmentForCoordinates(coords); !isAligned {
			sys.Events["OptimalAlignment"].AmountDays++
			sys.Events["OptimalAlignment"].DaysEvent = append(sys.Events["OptimalAlignment"].DaysEvent, dayChecked)
		}

	}
}
