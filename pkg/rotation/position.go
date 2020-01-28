package rotation

import pl "github.com/SolarSystem/pkg/planets"

type Position struct {
	Planet            pl.Planet
	ClockWisePosition float32
}

func New(p pl.Planet) *Position {
	pos := Position{
		p,
		0,
	}

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

// this calculates the angle between two points. this angle is the one of the sun.
func AngleBetweenPositions(p1 *Position, p2 *Position) float32 {
	if p1.ClockWisePosition > p2.ClockWisePosition {
		return p1.ClockWisePosition - p2.ClockWisePosition
	} else {
		return p2.ClockWisePosition - p1.ClockWisePosition
	}
}
