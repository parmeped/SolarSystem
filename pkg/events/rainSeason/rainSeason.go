package rainseason

import(
	
	m "math"
	sol "github.com/SolarSystem/pkg/system"
	pos "github.com/SolarSystem/pkg/position"
)

type RainSeason struct {
	Name string
	MaxPerimeter float32
}

// TODO: here, as in other methods, we should make a days calculator as to get amount of days on years. 
// GetRainSeasonsForYears returns how many rain seasons happen in {n} years, and when the most rainy day occurs.
func GetRainSeasonsForYears(years int, sys *sol.System) int {
	return GetRainSeasonsForDays((years * 365), sys)
}

// TODO: This logic is the same as the other events. This is a candidate for a generic, or at least a strategy.
// GetRainSeasonsForDays returns how many rain seasons happen in {n} days, and when the most rainy day occurs.
func GetRainSeasonsForDays(days int, sys *sol.System) int {
	cycleDays := pos.TimeToSystemCycle(sys.Positions[0], sys.Positions[1], sys.Positions[2])
	multiplier := days / int(cycleDays)
	daysRemaining := days % int(cycleDays)

	// Register event & check
	rainSeasonEvent := RainSeason{"RainSeason", 0}
	sys.AddCheck(rainSeasonEvent)
	sys.NewEvent(rainSeasonEvent.Name)

	rainSeasons, rainDays := getRainSeasonsForCycle(int(cycleDays), sys)
	rainSeasons = rainSeasons * multiplier

	// see if there's more events happening 
	for _, v := range rainDays {
		if v <= daysRemaining {
			rainSeasons++
		} else {
			break
		}
	}
	return rainSeasons
}

func getRainSeasonsForCycle(cycleDays int, sys *sol.System) (int, []int) {
	// Execute a complete cycle check on the system. 
	sol.RotateAndExecute(cycleDays, sys)	

	return sys.Events["RainSeason"].AmountDays, sys.Events["RainSeason"].DaysEvent
}

// DailyCheck function used for daily checks on a system after it rotates one day
func (rainy RainSeason) DailyCheck(sys *sol.System, dayChecked int) {	
	if isInside, coords := pos.ConvertToCartesianAndExecute(sys.Positions, checkSunInsideTriangle); isInside {
		sys.Events["RainSeason"].AmountDays++
		sys.Events["RainSeason"].DaysEvent = append(sys.Events["RainSeason"].DaysEvent, dayChecked)
		// here we should check for max perimeter, and set the peak day. 
		if peakRainDay, newPerimeter := checkForPeakRainDay(coords, rainy.MaxPerimeter); peakRainDay {
			sys.Events["RainSeason"].PeakDay = dayChecked
			rainy.MaxPerimeter = newPerimeter
		}		
	}
}


func checkForPeakRainDay(coords *[]pos.Coordinate, perimeter float32) (bool, float32) {
	c1, c2, c3 := (*coords)[0], (*coords)[1], (*coords)[2]
	var sideA, sideB, sideC float64
	sideA = m.Sqrt(m.Pow(float64(c2.X - c3.X), 2) + m.Pow(float64(c2.Y - c3.Y), 2))
	sideB = m.Sqrt(m.Pow(float64(c1.X - c3.X), 2) + m.Pow(float64(c1.Y - c3.Y), 2))
	sideC = m.Sqrt(m.Pow(float64(c1.X - c2.X), 2) + m.Pow(float64(c1.Y - c2.Y), 2))
	return (sideA + sideB + sideC) > float64(perimeter), float32(perimeter)
}

func checkSunInsideTriangle(coords *[]pos.Coordinate) (bool, *[]pos.Coordinate) {
	// sun coordinates
	sunCoords := pos.Coordinate{nil, 0, 0}
	
	b1 := Sign(sunCoords, (*coords)[0], (*coords)[1]) < 0
	b2 := Sign(sunCoords, (*coords)[1], (*coords)[2]) < 0
	b3 := Sign(sunCoords, (*coords)[2], (*coords)[0]) < 0
	return (b1 == b2) && (b2 == b3), coords
}

func Sign(c1, c2, c3 pos.Coordinate) float32 {
	return (c1.X - c3.X) * (c2.Y - c3.Y) - (c2.X - c3.X) * (c1.Y - c3.Y)
}

