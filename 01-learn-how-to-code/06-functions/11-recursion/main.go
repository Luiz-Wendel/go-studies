package main

import "fmt"

func main() {
	fmt.Println(recursiveFactorial(4))
	fmt.Println(loopFactorial(4))
}

func loopFactorial(number int) int {
	total := 1

	for ; number > 1; number-- {
		total *= number
	}

	return total
}

func recursiveFactorial(number int) int {
	// to avoid infinite loops we need to create a stop statement
	// a stop statement will end the recursive loop
	if number == 1 {
		return 1
	}

	// to create a recursive function we just need to call the same function inside its block scope
	return number * recursiveFactorial(number-1)
}
