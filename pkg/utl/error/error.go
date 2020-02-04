package error

import "fmt"

// HandleError super simple error handler to avoid panics
func HandleError(f string) {	
	if r := recover(); r != nil {
		fmt.Println("Recovered in", f, r)
	}
}