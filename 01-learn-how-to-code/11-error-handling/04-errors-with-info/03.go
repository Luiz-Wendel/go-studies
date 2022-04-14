package main

import (
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
		// `fmt.Errorf` formats the string to an `*errors.errorString`
		return 0, fmt.Errorf("square root of negative number: %v", f)
	}

	return 42, nil
}
