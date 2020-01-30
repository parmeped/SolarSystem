package main

import (
	"flag"
	"fmt"

	//h "github.com/SolarSystem/pkg/helpers"
	pos "github.com/SolarSystem/pkg/position"
	repo "github.com/SolarSystem/pkg/repository"
	sol "github.com/SolarSystem/pkg/system"
	"github.com/SolarSystem/pkg/utl/config"
)

func main() {

	cfgPath := flag.String("p", "conf.local.yaml", "Path to config file")
	flag.Parse()

	_, err := config.Load(*cfgPath)
	checkErr(err)

	DB := repo.New()
	sys := DB.SolarSystem


	var time1, time2, intersects = pos.GetTwoPointsIntersections(sys.Positions[1], sys.Positions[2])
	fmt.Printf("Amount intersects: %v , time1: %v, time2: %v \n", intersects, time1, time2)

	cycleDays := int(pos.TimeToSystemCycle(sys.Positions[0], sys.Positions[1], sys.Positions[2]))
	amountDroughs, daysDroughs := pos.GetDroughSeasonsForCycle(cycleDays, sys.Positions)
	fmt.Printf("amountDroughs: %v, daysDroughs: %v \n", amountDroughs, daysDroughs)
	droughsTotal := pos.GetDroughSeasonsForYears(0, sys.Positions)
	fmt.Printf("daysDroughs: %v \n", droughsTotal)

}

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func showPlanetsPositions(sol *sol.System) {
	fmt.Print("Reading positions...\n")
	for _, v := range sol.Positions {
		fmt.Printf("The planet %v is at the position %v \n", v.Planet.Name, v.ClockWisePosition)
	}
	fmt.Print("-------------------------------\n")
}

func showPlanetsData(sol *sol.System) {
	fmt.Print("Reading data...\n")
	for _, v := range sol.Positions {
		fmt.Printf("Planet: %v, Distance: %v, RotationSpeed: %v, TimeToCylce: %v \n",
			v.Planet.Name, v.Planet.Distance, v.Planet.Rotation_grades, v.Planet.TimeToCycle)
	}
	fmt.Print("-------------------------------\n")
}
