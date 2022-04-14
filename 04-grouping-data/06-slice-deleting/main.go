package main

import "fmt"

func main() {
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x)

	x = append(x, 77, 42, 532, 1234)
	fmt.Println(x)

	y := []int{2345, 342, 6354, 23}
	x = append(x, y...)
	fmt.Println(x)

	// using slicing and append to remove some values
	x = append(x[:2], x[4:]...)
	fmt.Println(x)
}
