/*
	● create a func with the identifier foo that;
		○ takes in a variadic parameter of type int;
		○ pass in a value of type []int into your func (unfurl the []int);
		○ returns the sum of all values of type int passed in.
	● create a func with the identifier bar that.
		○ takes in a parameter of type []int;
		○ returns the sum of all values of type int passed in.
*/

package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	fmt.Println(foo(numbers...))
	fmt.Println(bar(numbers))
}

func foo(numbers ...int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}

func bar(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}
