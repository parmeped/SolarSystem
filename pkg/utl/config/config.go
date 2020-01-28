package config

import (
	"fmt"
	"io/ioutil"

	repo "github.com/SolarSystem/pkg/repository"
	yaml "gopkg.in/yaml.v2"
)

// Load returns Configuration struct
func Load(path string) (*Configuration, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error reading config file, %s", err)
	}
	var cfg = new(Configuration)
	if err := yaml.Unmarshal(bytes, cfg); err != nil {
		return nil, fmt.Errorf("unable to decode into struct, %v", err)
	}
	cfg.DB.Planets = repo.GetPlanets()
	return cfg, nil
}

// Configuration holds data necessery for configuring application
type Configuration struct {
	DB        *Database
	Time_Vars *Time_Vars
}

type Time_Vars struct {
	Days_in_year       int     `yaml:"days_in_a_year,omitempty"`
	Years_to_calculate float32 `yaml:"years_to_calc,omitempty"`
	Starting_date      string  `yaml:"starting_date,omitempty"`
}

type Database struct {
	Planets *[]repo.Planet
}
