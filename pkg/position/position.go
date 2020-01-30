package position

import (
	"fmt"
	m "math"

	h "github.com/SolarSystem/pkg/helpers"
	pl "github.com/SolarSystem/pkg/planets"
)

type Position struct {
	Planet            pl.Planet
	ClockWisePosition float32
}

type intersections struct {
	timeToStart         float32
	timeToFirst         float32
	amountIntersections float32
	positionA           *Position
	positionB           *Position
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

// GetDroughSeasonsForYears returns the amount of droughs there's on a certain amount of years
// TODO: see if the hardcoded 365 value can be turned into a cfg call. What about leap years? This should receive a system, not positions
func GetDroughSeasonsForYears(years int, positions []*Position) int {
	return GetDroughSeasonsForDays((years * 365), positions)
}

// GetDroughSeasonsForDays returns the amount of droughs there's on a certain amount of days
func GetDroughSeasonsForDays(days int, positions []*Position) int {
	cycleDays := TimeToSystemCycle(positions[0], positions[1], positions[2])
	multiplier := days / int(cycleDays)
	daysRemaining := days % int(cycleDays)
	droughSeasons, droughDays := GetDroughSeasonsForCycle(int(cycleDays), positions)
	droughSeasons = droughSeasons * multiplier

	for _, v := range droughDays {
		if v <= daysRemaining {
			droughSeasons++
		} else {
			break
		}
	}
	return droughSeasons
}

// TODO: Unexport this. MAybe this could go inside the call to the other function, as a strategy or something.
// GetDroughSeasonsForCycle calculates how many times there's a Drough season on a cycle.
func GetDroughSeasonsForCycle(cycleDays int, positions []*Position) (int, []int) {
	// I know there's a least amount of time a couple of points can intersect. The check has to be for each pair of points
	fastestCycle, index, secondIndex := cycleDays, 0, 0

	// first get the two planets which are fastest to complete cycle
	for k, v := range positions {
		if int(v.Planet.TimeToCycle) < fastestCycle {
			index = k
			fastestCycle = int(v.Planet.TimeToCycle)
		}
	}
	fastestCycle = cycleDays

	for k, v := range positions {
		if k != index && int(v.Planet.TimeToCycle) < fastestCycle {
			secondIndex = k
			fastestCycle = int(v.Planet.TimeToCycle)
		}
	}

	timeStart, timeAny, amount := GetTwoPointsIntersections(positions[index], positions[secondIndex])
	var intersect = intersections{timeStart, timeAny, amount, positions[index], positions[secondIndex]}
	return checkForDroughs(intersect, cycleDays, positions)
}

// checks if there are droughs on a cycle. {amountOfDroughs, []daysOfDroughs}
func checkForDroughs(intersect intersections, cycleDays int, positions []*Position) (int, []int) {

	// the period starts on a drough, since all planets start on pos 0
	amountOfDroughs, days := 1, int(intersect.timeToFirst)
	positionToCheck, positionToCompare := &Position{}, &Position{}
	daysOfDroughs := []int{0} // TODO: this feels like a hack!

	// get the fastest planet
	if intersect.positionA.Planet.TimeToCycle > intersect.positionB.Planet.TimeToCycle {
		positionToCompare = intersect.positionB
	} else {
		positionToCompare = intersect.positionA
	}

	// get the position to check
	for _, v := range positions {
		if v != intersect.positionA && v != intersect.positionB {
			positionToCheck = v
		}
	}

	// TODO: check checkAlignmentForCoordinates on optimal alignment for
	var positionPlanetToCheck, positionPlanetToCompare int
	for i := days; i < cycleDays; {
		positionPlanetToCheck = GetPositionAtTime(&positionToCheck.Planet, i)
		positionPlanetToCompare = GetPositionAtTime(&positionToCompare.Planet, i)
		i = i + days
		if checkPositionsForDrough(positionPlanetToCheck, positionPlanetToCompare) {
			daysOfDroughs = append(daysOfDroughs, i)
			amountOfDroughs++
		}
	}
	return amountOfDroughs, daysOfDroughs
}

// drough check helper. Compares two positions to find a drough
func checkPositionsForDrough(positionToCheck, positionToCompare int) bool {
	result := positionToCheck - positionToCompare
	fmt.Printf("posCheck: %v, posCompare: %v, result: %v \n", positionToCheck, positionToCompare, result)

	if result < 0 {
		result = result * -1
	}

	if result == 180 || result == 0 {
		return true
	} else {
		return false
	}
}

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
	grades := float64(po.ClockWisePosition)
	return Coordinate{
		&pl,
		pl.Distance * float32((m.Cos(grades))),
		pl.Distance * float32((m.Sin(grades))),
	}
}
