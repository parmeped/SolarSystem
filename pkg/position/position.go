package position

import (
	m "math"

	h "github.com/SolarSystem/pkg/helpers"
	pl "github.com/SolarSystem/pkg/planets"
)

// Position is used to attach a planet to a position. 
type Position struct {
	Planet            pl.Planet
	ClockWisePosition float32
}

// TODO: Check if this can go
// Intersections 
type Intersections struct {
	TimeToStart         float32
	TimeToFirst         float32
	AmountIntersections float32
	PositionA           *Position
	PositionB           *Position
}

// TODO: check if the planet can go
// Coordinate struct for holding a planet coordinate.
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


// TODO: get that 360 on a config
// Moves a planet 1 day.
func Move(p *Position) {
	if p.Planet.Rotation_grades > 0 {
		p.ClockWisePosition = p.ClockWisePosition + p.Planet.Rotation_grades
		if p.ClockWisePosition >= 360 {
			p.ClockWisePosition = p.ClockWisePosition - 360
		}
	} else {
		if p.ClockWisePosition+p.Planet.Rotation_grades < 0 {
			p.ClockWisePosition = 360 - p.ClockWisePosition + p.Planet.Rotation_grades
		} else {
			p.ClockWisePosition = p.ClockWisePosition + p.Planet.Rotation_grades
		}
	}
}

// GetTwoPointsIntersections sees how many times two positions intersect. {timeStart, timeAnyPoint, amountIntersects}
func GetTwoPointsIntersections(p1, p2 *Position) (float32, float32, float32) {
	var amountIntersects float32 = 0
	var relativeSpeed = getRelativeSpeed(p1.Planet.Rotation_grades, p2.Planet.Rotation_grades)

	var timeToMeetAnyPoint = float32(360) / relativeSpeed
	var timeToMeetStart = timeToStartingPoint(p1, p2)

	amountIntersects = calculateMeetingPoints(timeToMeetStart, timeToMeetAnyPoint)
	return timeToMeetStart, timeToMeetAnyPoint, amountIntersects
}

// getRelativeSpeed returns the relative speed of two planets.
func getRelativeSpeed(speed1, speed2 float32) float32 {
	if speed1 > speed2 {
		return speed1 - speed2
	} 
	return speed2 - speed1	
}

// calculateMeetingPoints returns the amount of times they meet on a cycle
func calculateMeetingPoints(time1, time2 float32) float32 {
	if time1 > time2 {
		return time1 / time2
	} 
	return time2 / time1	
}

// TODO: get those 360 into a config
// GetPositionAtTime gets the position of a planet given a specific time.
func GetPositionAtTime(p *pl.Planet, days int) int {
	var travelledDistance = int(p.Rotation_grades) * days

	// negative value
	if travelledDistance < 0 {
		travelledDistance = travelledDistance + ((travelledDistance / -360) * 360)
		if travelledDistance != 0 {
			return 360 + travelledDistance
		} else {
			return travelledDistance
		}
	}

	// positive value
	if travelledDistance > 359 {
		return travelledDistance - ((travelledDistance / 360) * 360)
	} else {
		return travelledDistance
	}
}

// TimeToSystemCycle returns the time it takes for n positions to meet again at a starting point
func TimeToSystemCycle(p1, p2 *Position, positions ...*Position) float32 {
	result := timeToStartingPoint(p1, p2)

	for i := 0; i < len(positions); i++ {
		result = float32(h.LCM(int(result), int(positions[i].Planet.TimeToCycle)))
	}

	return result
}

// returns the amount of time it takes for two positions to meet at the starting point of a cycle
func timeToStartingPoint(p1, p2 *Position) float32 {
	return float32(h.LCM(int(p1.Planet.TimeToCycle), int(p2.Planet.TimeToCycle)))
}

// ConvertPolarToCartesian converts the position of a point and returns a cartesian c. The angular speed of the moving objects is in radians / t.
func ConvertPolarToCartesian(po *Position) Coordinate {
	pl := po.Planet
	// grades to radian conversion
	radians := float64(po.ClockWisePosition * 0.015708)
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
func ConvertToCartesianAndExecute(positions []*Position, fn CoordinatesBasedCheck) (bool, *[]Coordinate) {
	coordinates := []Coordinate{}
	for _, v := range positions {
		coordinates = append(coordinates, ConvertPolarToCartesian(v))
	}
	return fn(&coordinates)
}