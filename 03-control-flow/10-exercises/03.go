/*
	Create a for loop using this syntax
		● for condition { }
	Have it print out the years you have been alive.
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	year := 1994
	currentYear := time.Now().Year()

	for year <= currentYear {
		fmt.Println(year)
		year++
	}
}
