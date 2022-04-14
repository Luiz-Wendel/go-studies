package main

import "fmt"

func main() {
	x := []int{4, 5, 7, 8, 42}

	fmt.Println(x)
	fmt.Println(x[4])
	// prints every value from an index until the last
	fmt.Println(x[2:])
	// prints every value from the beggining until the second index (not included)
	fmt.Println(x[:2])
	// prints every value from the first index to the second index (not included)
	fmt.Println(x[1:4])
}
