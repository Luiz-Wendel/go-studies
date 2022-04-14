package main

import "fmt"

func main() {
	// the `make` built-in function takes the `type`, `length` and `capacity` to create a slice
	// the `capacity` is the size of the underlying `array`
	// this built-in function can be used to create map and channel as well
	// the following line is the same as creating an `array` of length 12 and slicing the first 10 values (`new([12]int)[0:10]`)
	x := make([]int, 10, 12)

	fmt.Println(x)
	fmt.Println(len(x))
	// prints the capacity
	fmt.Println(cap(x))

	// values can be assigned normally, respecting the slice length
	x[2] = 123

	// to increase the size we must use the `append` function
	x = append(x, 12)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	// if many values are appended to the fact that length is greater than the underlying array capacity,
	// Go will create another underlying array with double the capacity and throw previous values inside it
	x = append(x, 13)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 14)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
}
