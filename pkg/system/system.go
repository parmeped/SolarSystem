package system

import (
	p "github.com/SolarSystem/pkg/planets"
	pos "github.com/SolarSystem/pkg/position"
)

// Base system
type System struct {
	Positions   []*pos.Position
	DailyChecks []IDailyCheck
	Events      map[string]*Event
}

// Event struct
type Event struct {
	Name       string
	DaysEvent  []int
	AmountDays int
}

// IExecute is implemented by any event that wants to have a daily check performed.
type IDailyCheck interface {
	DailyCheck(sys *System, dayChecked int)
}

// Creates a new System with some planets
func New(planets []p.Planet) *System {
	sys := System{}
	for _, v := range planets {
		sys.Positions = append(sys.Positions, pos.New(v))
	}
	sys.DailyChecks = newChecks()
	sys.Events = make(map[string]*Event)
	return &sys
}

// Registers a new event on the system
func (sys *System) NewEvent(name string) {
	event := &Event{
		name,
		[]int{},
		0,
	}
	sys.Events[name] = event
}

// returns pointer to array of Checks
func newChecks() []IDailyCheck {
	return []IDailyCheck{}
}

// TODO: See if this can be implemented elsewhere
// AddEvent appends a new event to the system
func (sys *System) AddCheck(check IDailyCheck) {
	sys.DailyChecks = append(sys.DailyChecks, check)
}

// TODO: concurrency candidate.
// Rotates all planets on a system. Works like a charm.
func Rotate(days int, sys *System) {
	for i := 0; i < days; i++ {
		for _, v := range sys.Positions {
			pos.Move(v)
		}
		//performDroughCheck(sys) // first event check.
	}
}

// TODO: Does this make sense?
// RotateAndExecutes rotates {n} days and executes a function for each day. STARTS ON POSITION 0
func RotateAndExecute(days int, sys *System) {
	for i := 0; i < days; i++ {
		for _, v := range sys.Positions {
			pos.Move(v)
		}
		for _, v := range sys.DailyChecks {
			v.DailyCheck(sys, i)
		}
	}
}

// func performDroughCheck(sys *System) bool {
// 	var isDroughSeason = false
// 	if _, check := pos.AngleBetweenPositions(sys.Positions[0], sys.Positions[1]); check {
// 		isDroughSeason = e.CheckForDrough(sys.Positions[0], sys.Positions[2])
// 	}
// 	return isDroughSeason
// }
