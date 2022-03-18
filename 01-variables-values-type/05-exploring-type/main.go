package main

import "fmt"

// declares the variable with IDENTIFIER `y` with type `int`
var y = 42

// declares the variable with IDENTIFIER `z` with type `string` using double quotes
var z = "Shaken, not stirred"

// declares the variable with IDENTIFIER `x` with type `string` using back tick
var x = `James said,
"Shaken,

not stirred"`

func main() {
	fmt.Println(y)
	// using `Printf` to show the variable type
	fmt.Printf("%T\n", y)
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	// can't assign an integer to `z` as it is of type string
	// the following line would throw an error `cannot use 42 (type untyped int) as type string in assignment`
	// z = 42
	// fmt.Println(z)
	// fmt.Printf("%T\n", z)
}
