package position

import (
	m "math"

	h "github.com/SolarSystem/pkg/helpers"
	pl "github.com/SolarSystem/pkg/planets"
)

type Position struct {
	Planet            pl.Planet
	ClockWisePosition float32
}

type Intersections struct {
	TimeToStart         float32
	TimeToFirst         float32
	AmountIntersections float32
	PositionA           *Position
	PositionB           *Position
}

type Coordinate struct {
	Planet *pl.Planet
	X      float32
	Y      float32
}

func New(p pl.Planet) *Position {
	pos := Position{
		p,
		0,
	}
	return &pos
}

// TODO: The code is all over the place, and the functions have weird names. try and move a couple things out, and change func names. If there's time, see if it can be more generic

// Maybe a global config?
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

// TODO: Check, should be max 180°? When 180, it could be a drough season, if the third point is 180 too. I believe, since this is angled with the sun. I believe its malfunctioning, check this through
// Candidate for concurrency
// this calculates the angle between two points. this angle is the one of the sun.
func AngleBetweenPositions(p1 *Position, p2 *Position) (float32, bool) {
	//fmt.Printf("Planet: %v, Position: %v; Planet: %v, Position: %v \n", p1.Planet.Name, p1.ClockWisePosition, p2.Planet.Name, p2.ClockWisePosition)

	shouldCheckDrough := false
	var angle float32 = 0
	cwp1 := p1.ClockWisePosition
	cwp2 := p2.ClockWisePosition

	if cwp1 > 180 || cwp2 > 180 {
		if cwp1 > cwp2 {
			angle = (360 - cwp1) + cwp2
		} else {
			angle = (360 - cwp2) + cwp1
		}
	} else {
		if cwp1 > cwp2 {
			angle = cwp1 - cwp2
		} else {
			angle = cwp2 - cwp1
		}
	}

	if angle == 180 || angle == 0 {
		shouldCheckDrough = true
	}
	return angle, shouldCheckDrough
}

// DistanceBetweenPositions calculates the distance between two points.

// GetTwoPointsIntersections sees how many times two positions intersect. {timeStart, timeAnyPoint, amountIntersects}
// TODO: check if at given result time, there's an intersection with the remaining point. Further testing! this is just one check. there's more apparently
func GetTwoPointsIntersections(p1, p2 *Position) (float32, float32, float32) {
	var amountIntersects float32 = 0
	var relativeSpeed = getRelativeSpeed(p1.Planet.Rotation_grades, p2.Planet.Rotation_grades)

	var timeToMeetAnyPoint = float32(360) / relativeSpeed
	var timeToMeetStart = timeToStartingPoint(p1, p2)

	amountIntersects = calculateMeetingPoints(timeToMeetStart, timeToMeetAnyPoint)
	return timeToMeetStart, timeToMeetAnyPoint, amountIntersects
}

// gets relative speed of two planets.
func getRelativeSpeed(speed1, speed2 float32) float32 {
	if speed1 > speed2 {
		return speed1 - speed2
	} else {
		return speed2 - speed1
	}
}

func calculateMeetingPoints(time1, time2 float32) float32 {
	if time1 > time2 {
		return time1 / time2
	} else {
		return time2 / time1
	}
}

// GetPositionAtTime gets the position of a planet given a specific time. Seems to be working ok!
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

// get the time for n positions to complete a cycle.
func TimeToSystemCycle(p1, p2 *Position, positions ...*Position) float32 {
	result := timeToStartingPoint(p1, p2)

	for i := 0; i < len(positions); i++ {
		result = float32(h.LCM(int(result), int(positions[i].Planet.TimeToCycle)))
	}

	return result
}

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
