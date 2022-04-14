package main

import "fmt"

func main() {
	// x := type{values} // composite literal
	x := []int{4, 2, 32, 42}

	// prints the slice of `int` with identifier `x` created with composite literal
	fmt.Println(x)
	fmt.Println(len(x))
}
