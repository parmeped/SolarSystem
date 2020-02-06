package api

import (
	"log"
	"net/http"
	"os"
)

// SetupAndRunRouter runs the server
func SetupAndRunRouter() {

	http.HandleFunc("/ClimateForDay/", getClimateForDay)
	http.HandleFunc("/", handle)

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
}
