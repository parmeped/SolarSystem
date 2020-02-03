package main

import (
	"fmt"
	
	o "github.com/SolarSystem/pkg/events/optimalalignment"
	//e "github.com/SolarSystem/pkg/events"
	repo "github.com/SolarSystem/pkg/repository"
	sol "github.com/SolarSystem/pkg/system"
	//"github.com/SolarSystem/pkg/job"
	"github.com/SolarSystem/pkg/utl/config"
)

func main() {

	cfg := config.Load()	

	// Load initial planets
	DB := repo.New()
	DB.SolarSystem = sol.New(cfg.Planets, cfg)
	sys := DB.SolarSystem
	// Key: 1043 Value: &{1043 Tuesday, 12-Dec-23 22:07:08 -03 Óptima presión y temperatura. Ideal para la playa!}
	
	// job.CalculateModelForYears(1,sys,DB)
	// DB.ShowDaysModel()
	ok := o.SingleDayCheck(sys, 412)
	fmt.Printf("ok: %v \n", ok)	
	showSystemData(sys)

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
		fmt.Printf("Index: %v, Planet: %v, Distance: %v, RotationSpeed: %v, OrbitalPeriod: %v \n",
			k, v.Planet.Name, v.Planet.Distance, v.Planet.GradesPerDay, v.Planet.OrbitalPeriod)
	}
	fmt.Print("-------------------------------\n")
}

func showSystemData(sol *sol.System) {
	fmt.Print("Reading system data... \n")
	for _, v := range sol.Events {
		// fmt.Printf("Index: %v, Name: %v, Days event: %v, amountDays: %v \n",
		// 	k, v.Name, v.DaysEvent, v.AmountDays)
		fmt.Print(v)
		fmt.Print("\n")
	}
}


