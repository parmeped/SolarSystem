package planets

type Planet struct {
	Rotation_grades float32
	Distance        float32
	Name            string
	TimeToCycle     float32
}

// TODO: see if that 360 can be a cfg call
// Get all planets and set their time to complete 360 Cycle
func GetPlanets() []Planet {
	var planets = Planets_Array
	var i = 0
	for _, v := range planets {
		planets[i].TimeToCycle = 360 / v.Rotation_grades
		if planets[i].TimeToCycle < 0 {
			planets[i].TimeToCycle = planets[i].TimeToCycle * -1
		}
		i++
	}
	return planets
}
