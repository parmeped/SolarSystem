package planets

// Planet base struct
type Planet struct {
	GradesPerDay    int
	Distance        float32
	Name            string
	OrbitalPeriod   int
}

// New returns a new planet
func New(gradesPerDay, orbitalLength int, distance float32, name string) *Planet {
	planet := Planet{
		gradesPerDay,
		distance,
		name,
		0,
	}
	planet.OrbitalPeriod = calculateOrbitalPeriod(&planet, orbitalLength)
	return &planet
}


// GetPlanets get all planets and set their time to complete 1 OrbitalPeriod
func GetPlanets(orbitalLength int) []Planet {
	var planets = Planets_Array	
	for k, v := range planets {
		planets[k].OrbitalPeriod = calculateOrbitalPeriod(&v, orbitalLength)		
	}
	return planets
}

// The only one with Rotation_grades == 0 should be the sun. 
func calculateOrbitalPeriod(p *Planet, orbitalLength int) int {
	var time int = 0
	if p.GradesPerDay != 0 {		
		time = orbitalLength / int(p.GradesPerDay)
		if time < 0 {
			time = time * -1
		}		
	} 
	return time
}
