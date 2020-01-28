package events

import pos "github.com/SolarSystem/pkg/position"

// Seems to be working fine.
func CheckForDrough(p1 *pos.Position, p2 *pos.Position) bool {
	angle, _ := pos.AngleBetweenPositions(p1, p2)
	//fmt.Printf("Drough Check Angle: %v \n", angle)

	if angle == 180 || p1.ClockWisePosition == p2.ClockWisePosition {
		return true
	} else {
		return false
	}
}
