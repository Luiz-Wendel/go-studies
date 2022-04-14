package main

import "fmt"

func main() {
	// each variable will hold a different instance of the counter
	a := incrementor()
	b := incrementor()

	fmt.Println("a:", a())
	fmt.Println("a:", a())
	fmt.Println("a:", a())
	fmt.Println("b:", b())
	fmt.Println("b:", b())
}

func incrementor() func() int {
	// each call of the `incrementor` function will create a `counter` variable in different memory spaces
	var counter int

	return func() int {
		// as the function is inside the block scope of the `incrementor` function, it has access to the `counter` variable
		counter++

		return counter
	}
}
