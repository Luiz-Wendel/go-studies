package main

import "fmt"

func main() {
	sum := sum(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(sum)
}

// ? to declare a function that can receive any number of parameters you can use `...` operator before the parameter `type`
// ! if your function has more than one parameter, the `...` operator must be placed on the last parameter
func sum(numbers ...int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}
