package events

import (
	sol "github.com/SolarSystem/pkg/system"
	repo "github.com/SolarSystem/pkg/repository"
	config "github.com/SolarSystem/pkg/utl/config"
	"testing"
	"math/rand"
)


const (
	
)

func TestGetAmountPerEventForYearsFail(t *testing.T) {
	expectedResult := -1
	testName := "GetAmountEventForYearsFails"
	db := repo.New()
	cfg := config.Load()
	db.SolarSystem = sol.New(cfg.Planets, cfg)

	result := GetAmountPerEventForYears(1, db.SolarSystem, "testing")

	if result != expectedResult {
		t.Errorf("Error testing for %v. Want: %v, Got: %v", testName, expectedResult, result)
	}
}

func TestGetAmountPerEventForYearsSuccess(t *testing.T) {
	expectedResult := 0
	testName := "GetAmountEventForYearsSuccess"
	db := repo.New()
	cfg := config.Load()
	db.SolarSystem = sol.New(cfg.Planets, cfg)

	random := rand.Intn(2)

	result := GetAmountPerEventForYears(1, db.SolarSystem, getEvents()[random])

	if result < expectedResult {
		t.Errorf("Error testing for %v. Want higher than: %v, Got: %v",  testName, expectedResult, result)
	}
}

func TestGetAmountPerEventForDaysFail(t *testing.T) {
	expectedResult := -1
	testName := "GetAmountEventForDaysFail"
	db := repo.New()
	cfg := config.Load()
	db.SolarSystem = sol.New(cfg.Planets, cfg)

	result := GetAmountPerEventForDays(1, db.SolarSystem, "testing")

	if result != expectedResult {
		t.Errorf("Error testing for %v. Want: %v, Got: %v", testName, expectedResult, result)
	}
}

func TestGetAmountPerEventForDaysSuccess(t *testing.T) {
	expectedResult := 0
	testName := "GetAmountEventForDaysSuccess"
	db := repo.New()
	cfg := config.Load()
	db.SolarSystem = sol.New(cfg.Planets, cfg)

	random := rand.Intn(2)
	result := GetAmountPerEventForDays(1, db.SolarSystem, getEvents()[random])

	if result < expectedResult {
		t.Errorf("Error testing for %v. Want higher than: %v, Got: %v",  testName, expectedResult, result)
	}
}

func TestGetConditionForDayFail(t *testing.T) {
	db := repo.New()
	cfg := config.Load()
	db.SolarSystem = sol.New(cfg.Planets, cfg)

	result := GetConditionForDay(db.SolarSystem, -1)	
	testName := "TestGetConditionForDayFail"

	if result != nil {
		t.Errorf("Error testing for %v.",  testName)
	}
}

func TestGetConditionForDaySuccess(t *testing.T) {
	db := repo.New()
	cfg := config.Load()
	db.SolarSystem = sol.New(cfg.Planets, cfg)
	
	random := rand.Intn(1000)
	expectedResult := &sol.Day{Key:random}

	result := GetConditionForDay(db.SolarSystem, random)	
	testName := "TestGetConditionForDaySuccess"

	if result.Key != expectedResult.Key {
		t.Errorf("Error testing for %v. Want: %v, Got: %v",  testName, expectedResult, result)
	}
}

// workaround for "constant"
func getEvents() [3]string {
	return [...]string{"RainSeasaon", "DroughtSeason", "OptimalAlignment"}
}