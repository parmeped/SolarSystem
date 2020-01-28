package system

import (
	p "github.com/SolarSystem/pkg/planets"
	ro "github.com/SolarSystem/pkg/rotation"
)

type System struct {
	Positions []*ro.Position
}

func New(ps []p.Planet) *System {
	system := System{}
	for _, v := range ps {
		system.Positions = append(system.Positions, ro.New(v))
	}
	return &system
}

// TODO: concurrency candidate.
// Rotates all planets on a system. Works like a charm.
func Rotate(days int, sys *System) {
	for i := 0; i < days; i++ {
		for _, v := range sys.Positions {
			ro.Move(v)
		}
	}
}
