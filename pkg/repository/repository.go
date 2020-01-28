package repository

var planets = []Planet{
	Planet{
		5,
		false,
		1000,
		"Vulcano",
	},
	Planet{
		1,
		true,
		500,
		"Ferengi",
	},
	Planet{
		3,
		true,
		2000,
		"Betasoide",
	},
}

func GetPlanets() *[]Planet {
	return &planets
}

type Planet struct {
	Rotation_grades float32
	Clock_wise      bool
	Distance        float32
	Name            string
}
