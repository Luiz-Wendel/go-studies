package main

import "fmt"

func main() {
	foo()

	// as the `foo` function recovers from `panic`, this code executes normally
	fmt.Println("Returned normally from foo.")
}

func foo() {
	// built-in function `recover` only works inside deferred functions
	defer func() {
		// `recover` stops the panicking sequence by restoring normal execution
		// also retrieves the error value passed to the call of panic
		if r := recover(); r != nil {
			fmt.Println("Recovered in foo", r)
		}
	}()

	fmt.Println("Calling bar.")
	// calling panicking function `bar`
	bar(0)
	// this code does not execute, as `bar` function panics
	fmt.Println("Returned normally from bar.")
}

func bar(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		// panics with the value "4"
		panic(fmt.Sprintf("%v", i))
	}

	defer fmt.Println("Defer in bar", i)

	fmt.Println("Printing in bar", i)

	bar(i + 1)
}
