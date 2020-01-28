package position

import (
	pl "github.com/SolarSystem/pkg/planets"
)

type Position struct {
	Planet            pl.Planet
	ClockWisePosition float32
}

func New(p pl.Planet) *Position {
	pos := Position{
		p,
		0,
	}
	pos.Planet.TimeToCycle = 360 / pos.Planet.Rotation_grades
	return &pos
}

// TODO: Logic here should be something that checks for circle completition, thus restarting.
// Maybe a global config? We can momentarily assume the rotation_grades property COULD just be negative, indicating counterclock movement.
// this can be improved if we multiply the amount of grades it has to move. days * grades, should give amount.
// Moves a planet 1 day.
func Move(p *Position) {
	if p.Planet.Clock_wise {
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
