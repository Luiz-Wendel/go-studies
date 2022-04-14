package main

import "fmt"

func main() {
	x := []int{4, 5, 7, 8, 42}

	fmt.Println(x[0])
	fmt.Println(x[4])

	// the `range` keyword can be used to iterate over the slice, returning the `index` and the `value`
	for index, value := range x {
		fmt.Println(index, value)
	}
}
