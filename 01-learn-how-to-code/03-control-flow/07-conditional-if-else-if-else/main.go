package main

import "fmt"

func main() {
	x := 42

	// to add more conditions the `else` keyword can be used
	if x == 40 {
		fmt.Println("x is 40")
	} else if x == 41 {
		fmt.Println("x is 41")
	} else if x == 42 {
		fmt.Println("x is 42")
	} else {
		fmt.Println("x value is out of range")
	}
}
