package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	repo "github.com/SolarSystem/pkg/repository"
	service "github.com/SolarSystem/pkg/service"
	sol "github.com/SolarSystem/pkg/system"
	config "github.com/SolarSystem/pkg/utl/config"
)

const (
	processingError      = "There was an error processing the request"
	noParameterError     = "Error, no parameter was provided"
	dayGreaterThan0Error = "Error, day queried must be equal to or greater than 0"
	nothingToSee         = "Nothing to see here! Move along..."
	everythingOk         = "Found object and returning"
)

type response struct {
	StatusCode int
	Message    string
	Object     interface{}
}

// Default
func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, nothingToSee)
}

// GetClimateForDay handler
func getClimateForDay(w http.ResponseWriter, r *http.Request) {
	request := r.URL.Query()["Day"]
	response := response{}

	if request[0] == "" {
		response.StatusCode = http.StatusBadRequest
		response.Message = noParameterError
		marshallAndWrite(w, response)
		return
	}

	if day, err := strconv.ParseInt(request[0], 0, 0); err != nil {
		response.StatusCode = http.StatusInternalServerError
		response.Message = processingError
		marshallAndWrite(w, response)
		return
	} else {
		if day < 0 {
			response.StatusCode = http.StatusBadRequest
			response.Message = dayGreaterThan0Error
			marshallAndWrite(w, response)
			return
		}

		cfg := config.Load()

		// Load initial planets
		DB := repo.New()
		DB.SolarSystem = sol.New(cfg.Planets, cfg)

		service := service.New(DB)
		response.Object = service.GetClimateForDay(int(day))

		if response.Object != nil {
			response.StatusCode = http.StatusOK
			response.Message = everythingOk
			marshallAndWrite(w, response)
		} else {
			response.StatusCode = http.StatusInternalServerError
			response.Message = processingError
			marshallAndWrite(w, response)
		}
	}
}

func marshallAndWrite(w http.ResponseWriter, res response) {
	js, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
