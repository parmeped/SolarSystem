package events

// Event struct
type Event struct {
	key        string
	DaysEvent  []int
	AmountDays int
}

// returns pointer to event
func New(name string) *Event {
	return &Event{
		name,
		[]int{},
		0,
	}
}

// returns pointer to array of events
func NewEvents() *[]Event {
	return &[]Event{}
}
