package system

import (
	e "github.com/SolarSystem/pkg/events"
	p "github.com/SolarSystem/pkg/planets"
	pos "github.com/SolarSystem/pkg/position"
)

type System struct {
	Positions []*pos.Position
}

func New(ps []p.Planet) *System {
	system := System{}
	for _, v := range ps {
		system.Positions = append(system.Positions, pos.New(v))
	}
	return &system
}

// TODO: concurrency candidate.
// Rotates all planets on a system. Works like a charm.
func Rotate(days int, sys *System) {
	for i := 0; i < days; i++ {
		for _, v := range sys.Positions {
			pos.Move(v)
		}
		performDroughCheck(sys) // first event check.
	}
}

func performDroughCheck(sys *System) bool {
	var isDroughSeason = false
	if _, check := pos.AngleBetweenPositions(sys.Positions[0], sys.Positions[1]); check {
		isDroughSeason = e.CheckForDrough(sys.Positions[0], sys.Positions[2])
	}
	return isDroughSeason
}
