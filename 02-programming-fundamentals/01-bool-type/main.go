package main

import "fmt"

// declares the variable with identifier `x` as type `bool` (boolean |-> true/false)
var x bool

func main() {
	// prints the zero value of a boolean
	fmt.Println(x)
	x = true
	// prints the new value of `x`
	fmt.Println(x)

	a := 10
	b := 42

	// compares if value of `a` equals (double `=`) value of `b`
	// comparison returns a boolean
	fmt.Println(a == b)
}

/*
	Comparison operators:
		==    equal
		!=    not equal
		<     less
		<=    less or equal
		>     greater
		>=    greater or equal

	Source: https://go.dev/ref/spec#Comparison_operators
*/
