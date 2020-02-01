package main

import (
	"flag"
	"fmt"
	"time"

	//h "github.com/SolarSystem/pkg/helpers"

	o "github.com/SolarSystem/pkg/events/optimalalignment"
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
	timeStamp()

	days := o.GetOptimalAlignmentsForYears(10, sys)
	fmt.Printf("days %v \n", days)
	
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
	for k, v := range sol.Positions {
		fmt.Printf("Index: %v, Planet: %v, Distance: %v, RotationSpeed: %v, TimeToCylce: %v \n",
			k, v.Planet.Name, v.Planet.Distance, v.Planet.Rotation_grades, v.Planet.TimeToCycle)
	}
	fmt.Print("-------------------------------\n")
}

func showSystemData(sol *sol.System) {
	fmt.Print("Reading system data... \n")
	for k, v := range sol.Events {
		fmt.Printf("Index: %v, Name: %v, Days event: %v, amountDays: %v \n",
			k, v.Name, v.DaysEvent, v.AmountDays)
	}
}

func timeStamp() {
	fmt.Print("////////////////   Program starting... ////////////////// \n")
	fmt.Printf("Time: %v \n", time.Now())
	fmt.Print("-------------------------------\n")

}
