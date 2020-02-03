package repository

import (
	"fmt"
	"sort"
	sol "github.com/SolarSystem/pkg/system"	
)

// Database wich contains the SolarSystem
type Database struct {
	SolarSystem *sol.System
	Days		map[int]*sol.Day
}

// New Returns a pointer to a MockDb
func New() *Database {		
	db := Database{}
	db.Days = make(map[int]*sol.Day)
	return &db
}

// AddDaysModel add an array of day pointers to the database
func (db *Database) AddDaysModel(days []*sol.Day) {
	for _, v := range days {
		db.Days[v.Key] = v
	}	
}

// ShowDaysModel shows all days on the database
func (db *Database) ShowDaysModel() {	
	var keys []int
	for k := range db.Days {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		fmt.Println("Key:", k, "Value:", db.Days[k])
		fmt.Print("\n")
	}
}