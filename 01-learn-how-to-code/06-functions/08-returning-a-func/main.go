package main

import "fmt"

func main() {
	// the varaible will receive a function
	f := bar()

	// to check the variable type
	fmt.Printf("%T\n", f)
	// to execute a function
	fmt.Println(f())
}

// `func() int` is the `return type` of the `bar` function
func bar() func() int {
	return func() int {
		return 451
	}
}
