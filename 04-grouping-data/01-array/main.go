package main

import "fmt"

func main() {
	// to declare an array we must put the brackets before the type of the variable
	// inside the brackets we must define the length of the array, or else the array will have length `0`
	var x [5]int

	fmt.Println(x)
	// to assign a value in some position of the array we must use the following syntax:
	x[3] = 42
	fmt.Println(x)
	// arrays are zero indexed, so the first position is `0` and the last is `length - 1`
	// the following line would throw an error `invalid array index 5 (out of bounds for 5-element array)`
	// x[5] = 2

	// we can check the array length with the `len()` function
	fmt.Println(len(x))
}

/*
	? There are major differences between the ways arrays work in Go and C. In Go,
		? Arrays are values. Assigning one array to another copies all the elements.
		? In particular, if you pass an array to a function, it will receive a copy of the array, not a pointer to it.
		? The size of an array is part of its type. The types [10]int and [20]int are distinct.

	? The value property can be useful but also expensive;
	? if you want C-like behavior and efficiency, you can pass a pointer to the array.

	? But even this style isn't idiomatic Go.
	! Use slices instead.

	Source: https://go.dev/doc/effective_go#arrays
*/
