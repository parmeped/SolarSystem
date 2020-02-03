package system

// Day base struct
type Day struct {
	Key		  int
	Date 	  string
	Condition string
}

const (
	// Rainy day
	Rainy 	= "Lluvioso"
	// PeakRain day
	PeakRain = "Hoy es el día de máxima lluvia!"	
	// Optimal day
	Optimal = "Óptima presión y temperatura. Ideal para la playa!"
	// Drought day
	Drought = "Sequía"
	// Normal day
	Normal 	= "Condiciones climáticas normales"
)