package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	// api "github.com/SolarSystem/pkg/api"
	// repo "github.com/SolarSystem/pkg/repository"
	// service "github.com/SolarSystem/pkg/service"
	// sol "github.com/SolarSystem/pkg/system"
	// config "github.com/SolarSystem/pkg/utl/config"
)

func main() {
	http.HandleFunc("/", indexHandler)

	// [START setting_port]
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
	// [END setting_port]
}

// [END main_func]

// [START indexHandler]

// indexHandler responds to requests with our greeting.
func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprint(w, "Hello, World!")
}

// func main() {

// 	cfg := config.Load()

// 	// Load initial planets
// 	DB := repo.New()
// 	DB.SolarSystem = sol.New(cfg.Planets, cfg)

// 	service := service.New(DB)
// 	r := api.SetupRouter(service)
// 	r.Run()

// }
