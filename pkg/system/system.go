package system

import (
	p "github.com/SolarSystem/pkg/planets"
	pos "github.com/SolarSystem/pkg/position"
	c "github.com/SolarSystem/pkg/utl/config"
)

// TODO: remove sun from here for consistency. Check where it's used
// Base system
type System struct {
	Positions      []*pos.Position
	Events         map[string]*Event
	SunCoordinates *pos.Coordinate // what if it has multiple suns?
	Cfg            *c.Configuration
	TimeToCycle    int
}

// Event struct
type Event struct {
	Name            string
	DaysEvent       []int
	AmountDays      int
	PeakDay         int
	MaxPerimeter    float32
	Implementations IEvent
	DailyCheck      bool
}

// IEvent is implemented by any event that wants to have a daily check performed.
type IEvent interface {
	DailyCheck(sys *System, dayChecked int)
	GetEventPerCycle(cycleDays int, sys *System) (int, []int)
}

// New Creates a new System with some planets
func New(planets []p.Planet, cf *c.Configuration) *System {
	sys := System{}
	for _, v := range planets {
		sys.Positions = append(sys.Positions, pos.New(v))
	}
	//sys.DailyChecks = newChecks()
	sys.Events = make(map[string]*Event)
	sys.Cfg = cf
	theSun := p.New(0, cf.Orbit, 0, "Sun")
	sys.SunCoordinates = &pos.Coordinate{theSun, 0, 0}
	return &sys
}

// NewEvent registers a new event on the system
func (sys *System) NewEvent(name string, implementation IEvent, dailyCheck bool) {
	event := &Event{
		name,
		[]int{},
		0,
		0,
		0,
		implementation,
		dailyCheck,
	}
	sys.Events[name] = event
}

// returns pointer to array of Checks
func newChecks() []IEvent {
	return []IEvent{}
}

// RotateAndExecute rotates {n} days and executes a function for each day. STARTS ON POSITION 0
// TODO: this gets called per event instead of only once. Change that!
func RotateAndExecute(days int, sys *System, event string) {
	for i := 0; i < days; i++ {
		sys.Events[event].Implementations.DailyCheck(sys, i)
		for _, v := range sys.Positions {
			pos.Move(v)
		}
	}
}
