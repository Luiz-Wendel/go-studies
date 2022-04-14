package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4}

	fmt.Println(sum(numbers...))

	fmt.Println(sumEven(sum, numbers...))
}

func sum(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

// passing a function (callback function) as a parameter `identifier func([parameters]) [return type]`
func sumEven(f func(numbers ...int) int, numbers ...int) int {
	var evenNumbers []int

	for _, number := range numbers {
		if number%2 == 0 {
			evenNumbers = append(evenNumbers, number)
		}
	}

	// the identifier of the callback will hold the function which you can call using parentheses
	return f(evenNumbers...)
}
