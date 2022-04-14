package main

import "fmt"

var (
	y string
	z int
)

func main() {
	fmt.Println("Printing the `zero value` of 'y':")
	// the following line will print an empty string
	fmt.Println(y)

	y = "Bond, James Bond"

	fmt.Println("Printing the new value of 'y':")
	fmt.Println(y)

	fmt.Println("Printing the `zero value` of 'z':")
	fmt.Println(z)
}

/*
	ZERO VALUES list:
		- Booleans: false
		- Integers: 0
		- Floats: 0.0
		- Strings: ""
		- Others: nil
			|-> (Pointers, Functions, Interfaces, Slices, Channels, Maps)
*/
