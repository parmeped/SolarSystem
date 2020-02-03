package main

import (

	sol "github.com/SolarSystem/pkg/system"
	"github.com/SolarSystem/pkg/utl/config"
	"github.com/SolarSystem/pkg/service"
	"github.com/SolarSystem/pkg/api"
	repo "github.com/SolarSystem/pkg/repository"
)

func main() {

	cfg := config.Load()	

	// Load initial planets
	DB := repo.New()
	DB.SolarSystem = sol.New(cfg.Planets, cfg)	
	
	service := service.New(DB)
	r := api.SetupRouter(service)

	r.Run()
}