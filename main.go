package main

import (
	"fmt"
	"net/http"
)

func main() {

	// cfg := config.Load()

	// // Load initial planets
	// DB := repo.New()
	// DB.SolarSystem = sol.New(cfg.Planets, cfg)

	// service := service.New(DB)
	// r := api.SetupRouter(service)
	// r.Run()
	http.HandleFunc("/", handle)

}

func handle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprint(w, "Hello world!")
}
