/*
	A “callback” is when we pass a func into a func as an argument.

	For this exercise,
		● pass a func into a func as an argument
*/

package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 42}

	fmt.Println(calculator(sum, numbers))
	fmt.Println(calculator(mult, numbers))
}

func calculator(f func(numbers ...int) int, numbers []int) int {
	return f(numbers...)
}

func sum(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

func mult(numbers ...int) int {
	total := 1

	for _, number := range numbers {
		total *= number
	}

	return total
}
