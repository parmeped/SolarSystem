package main

import (
	"flag"
	"fmt"

	sol "github.com/SolarSystem/pkg/system"
	"github.com/SolarSystem/pkg/utl/config"
)

func main() {

	cfgPath := flag.String("p", "conf.local.yaml", "Path to config file")
	flag.Parse()

	cfg, err := config.Load(*cfgPath)
	checkErr(err)
	sys := cfg.DB.SolarSystem

	for i := 0; i < 1500; i++ {
		sol.Rotate(i, sys)
	}
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
