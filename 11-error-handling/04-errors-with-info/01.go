package main

import (
	"errors"
	"fmt"
)

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
		// if a negative number is passed, returns an error with info `square root of negative number`
		return 0, errors.New("square root of negative number")
	}

	return 42, nil
}
