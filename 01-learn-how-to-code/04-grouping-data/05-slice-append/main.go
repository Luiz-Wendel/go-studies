package main

import "fmt"

func main() {
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x)

	// the `append` function receives a `slice` in the first parameter and a variadic number of elements as the second parameter
	// it will return a new `slice` with all values together
	x = append(x, 77, 42, 532, 1234)
	fmt.Println(x)

	// we can spread the values of another `slice` using the `...` after the `slice` identifier
	y := []int{2345, 342, 6354, 23}
	x = append(x, y...)
	fmt.Println(x)
}
