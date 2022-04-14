package main

import (
	"errors"
	"fmt"
)

// creates an error variable
var ErrorSqrt = errors.New("square root of negative number")

func main() {
	_, error := sqrt(-42)

	if error != nil {
		fmt.Println(error)
		fmt.Printf("%T\n", error)

		fmt.Println(error.Error())
		fmt.Printf("%T\n", error.Error())
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		// using the error variable
		return 0, ErrorSqrt
	}

	return 42, nil
}
