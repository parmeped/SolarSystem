package planets

type Planet struct {
	Rotation_grades float32
	Distance        float32
	Name            string
	TimeToCycle     float32
}

func New(rotation, distance float32, name string) *Planet {
	planet := Planet{
		rotation,
		distance,
		name,
		0,
	}
	planet.TimeToCycle = calculateTimeToCycle(&planet)
	return &planet
}

// TODO: see if that 360 can be a cfg call
// Get all planets and set their time to complete 360 Cycle
func GetPlanets() []Planet {
	var planets = Planets_Array
	var i = 0
	for _, v := range planets {
		v.TimeToCycle = calculateTimeToCycle(&v)
		i++
	}
	return planets
}

// TODO: this would break everything if any other planet aside from the sun has a timeToCycle = 0
func calculateTimeToCycle(p *Planet) float32 {
	if p.Rotation_grades == 0 {
		return 0
	} else {
		time := 360 / p.Rotation_grades
		if time < 0 {
			time = time * -1
		}
		return time
	}
}
