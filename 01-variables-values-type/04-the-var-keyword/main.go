package main

import "fmt"

// declare the `y` variable in the global scope
var y = 84

// declare the `z` variable with type `int` and value `0`
var z int

// the following line would throw an error `syntax error: non-declaration statement outside function body`
// x := 42
// short declaration operator can only be used inside function scope

func main() {
	x := 42

	// can access `x` because it was declared inside function scope
	fmt.Println("x:", x)
	// can access `y` because it is in the global scope
	fmt.Println("first y:", y)

	foo()

	fmt.Println("z:", z)
}

func foo() {
	// also can access `y`
	fmt.Println("second y:", y)
	// the following line would throw an error `undefined: x` because `x` scope is not declared inside this function scope
	// fmt.Println("another x:", x)
}
