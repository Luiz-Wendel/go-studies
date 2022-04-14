package main

import "fmt"

type sqrtError struct {
	value float64
	err   error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("square root error: %v (%v)", se.value, se.err)
}

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
		se := fmt.Errorf("square root of negative number: %v", f)
		return 0, sqrtError{f, se}
	}

	return 42, nil
}
