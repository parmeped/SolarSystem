package main

import (
	"flag"
	"fmt"

	//h "github.com/SolarSystem/pkg/helpers"
	pos "github.com/SolarSystem/pkg/position"
	sol "github.com/SolarSystem/pkg/system"
	repo "github.com/SolarSystem/pkg/repository"
	"github.com/SolarSystem/pkg/utl/config"
)

func main() {

	cfgPath := flag.String("p", "conf.local.yaml", "Path to config file")
	flag.Parse()

	_, err := config.Load(*cfgPath)
	checkErr(err)
	
	DB := repo.New()
	sys := DB.SolarSystem

	var intersects = pos.IntersectsChecks(sys.Positions[1], sys.Positions[2], 180)
	fmt.Printf("Amount intersects: %v \n", intersects)
	sol.Rotate(180, sys)
	showPlanetsPositions(sys)	

	var distance = pos.GetPositionPointTime(&sys.Positions[0].Planet, 1500)

	fmt.Printf("Distance: %v", distance)

	//fmt.Println(h.LCM(12, 80))

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
