package main

import "fmt"

func main() {
	// `if` block scope only executes if the condition resolves to `true`
	if 42 == 42 {
		fmt.Println("01")
	}

	// the `!` can be used to toggle the boolean value
	// `!true` becomes `false` and `!false` becomes `true`
	if !(42 == 42) {
		fmt.Println("02")
	}

	// when we want more than one statement in the same line, in Go, we need to use the `;` at the end of the first statement
	// in the following line we declare a variable that will only be available for the `if` scope
	if x := 42; x == 42 {
		fmt.Println("03")
	}
	// the following line would throw an error `undefined: x`
	// fmt.Println(x)
}
