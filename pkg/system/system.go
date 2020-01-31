package system

import (
	e "github.com/SolarSystem/pkg/events"
	p "github.com/SolarSystem/pkg/planets"
	pos "github.com/SolarSystem/pkg/position"
)

// Base system
type System struct {
	Positions []*pos.Position
	Events    *[]e.Event
}

// IExecute is implemented by any event that wants to have a daily check performed.
type IExecute interface {
	DailyCheck(sys *System)
}

// Creates a new System with some planets
func New(planets []p.Planet) *System {
	sys := System{}
	for _, v := range planets {
		sys.Positions = append(sys.Positions, pos.New(v))
	}
	sys.Events = e.NewEvents()
	return &sys
}

// TODO: See if this can be implemented elsewhere
// AddEvent appends a new event to the system
func (sys *System) AddEvent(event *e.Event) {
	*sys.Events = append(*sys.Events, (*event))
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

// RotateAndExecutes rotates {n} days and executes a function for each day.
func RotateAndExecute(days int, sys *System, fn []IExecute) {
	for i := 0; i < days; i++ {
		for _, v := range sys.Positions {
			pos.Move(v)
		}
		for _, v := range fn {
			v.DailyCheck(sys)
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
