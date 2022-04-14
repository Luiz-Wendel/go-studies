/*
	● Create a func which returns a func;
	● assign the returned func to a variable;
	● call the returned func.
*/

package main

import "fmt"

func main() {
	first := decrementor(100)
	second := decrementor(10)

	fmt.Println(first())
	fmt.Println(first())
	fmt.Println(first())
	fmt.Println(second())
	fmt.Println(second())
}

func decrementor(counter int) func() int {
	return func() int {
		counter--

		return counter
	}
}
