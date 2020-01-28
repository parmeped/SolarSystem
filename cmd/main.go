package main

import (
	"flag"
	"fmt"

	h "github.com/SolarSystem/pkg/helpers"
	sol "github.com/SolarSystem/pkg/system"
	"github.com/SolarSystem/pkg/utl/config"
)

func main() {

	cfgPath := flag.String("p", "conf.local.yaml", "Path to config file")
	flag.Parse()

	_, err := config.Load(*cfgPath)
	checkErr(err)
	//sys := cfg.DB.SolarSystem

	fmt.Println(h.LCM(12, 80))

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
