package main

import "fmt"

func main() {
	fmt.Println("2 + 3 =", sum(2, 3))
	fmt.Println("4 + 7 =", sum(4, 7))
	fmt.Println("5 + 9 =", sum(5, 9))
}

// declared the `sum` function which will be tested
func sum(numbers ...int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}
