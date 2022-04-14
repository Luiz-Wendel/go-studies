package main

import "fmt"

func main() {
	// anonymous self executing function
	func() {
		fmt.Println("Anonymous function called")
	}()

	// anonymous function with parameters
	func(x int) {
		fmt.Println("Anonymous function called with", x)
	}(42)

	// anonymous function with return
	result := func(x int) int {
		return x
	}(42)

	fmt.Println(result)
}
