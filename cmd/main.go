package main

import (
	"flag"
	"fmt"

	"github.com/SolarSystem/pkg/utl/config"
)

func main() {

	cfgPath := flag.String("p", "conf.local.yaml", "Path to config file")
	flag.Parse()

	cfg, err := config.Load(*cfgPath)
	checkErr(err)

	fmt.Printf("The distance from the sun for planet %v is %v", cfg.DB.Planets.get(0), cfg.DB.Planets.get(0))
}

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
