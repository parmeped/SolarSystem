package position

import (
	m "math"

	h "github.com/SolarSystem/pkg/helpers"
	pl "github.com/SolarSystem/pkg/planets"
	c "github.com/SolarSystem/pkg/utl/config"
)

// Position is used to attach a planet to a position.
type Position struct {
	Planet            pl.Planet
	ClockWisePosition int
}

// Coordinate struct for holding a planet coordinate.
// TODO: [Improvement] Check if the planet can go
type Coordinate struct {
	Planet *pl.Planet
	X      float32
	Y      float32
}

// CoordinatesBasedCheck is used pass a func to the ConvertToCartesianAndExecute() and later exectue it.
type CoordinatesBasedCheck func(coords *[]Coordinate) (bool, *[]Coordinate)

// New returns a pointer to a position
func New(p pl.Planet) *Position {
	pos := Position{
		p,
		0,
	}
	return &pos
}

// Move moves a planet 1 day.
func Move(p *Position) {
	orbit := c.GetOrbit()
	if p.Planet.GradesPerDay > 0 {
		p.ClockWisePosition = p.ClockWisePosition + p.Planet.GradesPerDay
		if p.ClockWisePosition >= orbit {
			p.ClockWisePosition = p.ClockWisePosition - orbit
		}
	} else {
		if p.ClockWisePosition+p.Planet.GradesPerDay < 0 {
			p.ClockWisePosition = orbit - p.ClockWisePosition + p.Planet.GradesPerDay
		} else {
			p.ClockWisePosition = p.ClockWisePosition + p.Planet.GradesPerDay
		}
	}
}

// GetTwoPointsIntersections sees how many times two positions intersect. {timeStart, timeAnyPoint, amountIntersects}
func GetTwoPointsIntersections(p1, p2 *Position) (int, int, int) {
	amountIntersects := 0
	var relativeSpeed = getRelativeSpeed(p1.Planet.GradesPerDay, p2.Planet.GradesPerDay)

	var timeToMeetAnyPoint = int(c.GetOrbit() / relativeSpeed)
	var timeToMeetStart = timeToStartingPoint(p1, p2)

	amountIntersects = calculateMeetingPoints(timeToMeetStart, timeToMeetAnyPoint)
	return timeToMeetStart, timeToMeetAnyPoint, amountIntersects
}

// getRelativeSpeed returns the relative speed of two planets.
func getRelativeSpeed(speed1, speed2 int) int {
	if speed1 > speed2 {
		return speed1 - speed2
	}
	return speed2 - speed1
}

// calculateMeetingPoints returns the amount of times they meet on a cycle
func calculateMeetingPoints(time1, time2 int) int {
	if time1 > time2 {
		return time1 / time2
	}
	return time2 / time1
}

// GetPositionAtTime gets the position of a planet given a specific time. Works fine
func GetPositionAtTime(p *pl.Planet, days int) int {
	var travelledDistance = int(p.GradesPerDay) * days
	orbit := c.GetOrbit()

	pos := travelledDistance % orbit
	if pos < 0 {
		return pos + orbit
	}
	return pos
}

// TimeToSystemCycle returns the time it takes for n positions to meet again at a starting point
func TimeToSystemCycle(p1, p2 *Position, positions ...*Position) int {
	result := timeToStartingPoint(p1, p2)

	for i := 0; i < len(positions); i++ {
		result = h.LCM(result, positions[i].Planet.OrbitalPeriod)
	}

	return result
}

// returns the amount of time it takes for two positions to meet at the starting point of a cycle
func timeToStartingPoint(p1, p2 *Position) int {
	return h.LCM(p1.Planet.OrbitalPeriod, p2.Planet.OrbitalPeriod)
}

// ConvertPolarToCartesian converts the position of a point and returns a cartesian c. The angular speed of the moving objects is in radians / t.
func ConvertPolarToCartesian(po *Position) Coordinate {
	pl := po.Planet
	// grades to radian conversion
	radians := float64(float32(po.ClockWisePosition) * 0.015708)
	x, y := pl.Distance*float32((m.Cos(radians))), pl.Distance*float32((m.Sin(radians)))
	x, y = float32(m.Floor(float64(x)*100)/100), float32(m.Floor(float64(y)*100)/100)
	return Coordinate{
		&pl,
		x,
		y,
	}
}

// ConvertPolarSliceToCartesian converts a slice of positions to cartesian coordinates
func ConvertPolarSliceToCartesian(positions []*Position) *[]Coordinate {
	coords := &[]Coordinate{}
	for _, v := range positions {
		*coords = append(*coords, ConvertPolarToCartesian(v))
	}
	return coords
}

// ConvertToCartesianAndExecute converts polar positions to cartesian and then executes a func
// This was used before when CheckAlignmentForCoordinates was implemented on OptimalAlignment. Since it's kind of the same, I've left it here.
func ConvertToCartesianAndExecute(positions []*Position, fn CoordinatesBasedCheck) (bool, *[]Coordinate) {
	coordinates := []Coordinate{}
	for _, v := range positions {
		coordinates = append(coordinates, ConvertPolarToCartesian(v))
	}
	return fn(&coordinates)
}

// CheckAlignmentForCoordinates checks if {n} coordinates are aligned.
func CheckAlignmentForCoordinates(coords *[]Coordinate) (bool, *[]Coordinate) {
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
		slopeA, slopeB = m.Floor(slopeA*10)/10, m.Floor(slopeB*10)/10

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
