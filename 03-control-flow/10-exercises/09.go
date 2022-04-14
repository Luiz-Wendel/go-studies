/*
	Create a program that uses a switch statement with the switch expression specified as a variable of TYPE string with the IDENTIFIER “favSport”.
*/

package main

import "fmt"

func main() {
	favSport := "Some sport"

	switch favSport {
	case "Tennis":
		fmt.Println("Your favorite sport is Tennis")
	case "Volleyball":
		fmt.Println("Your favorite sport is Volleyball")
	case "Ping Pong":
		fmt.Println("Your favorite sport is Ping Pong")
	default:
		fmt.Println("You have no favorite sport")
	}
}
