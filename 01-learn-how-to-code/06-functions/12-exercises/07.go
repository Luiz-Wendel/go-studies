/*
	● Assign a func to a variable, then call that func.
*/

package main

import "fmt"

func main() {
	sum := func(numbers ...int) {
		total := 0

		for _, number := range numbers {
			total += number
		}

		fmt.Println(total)
	}

	sum(42, 1, 5, 2, 3, 4)
}
