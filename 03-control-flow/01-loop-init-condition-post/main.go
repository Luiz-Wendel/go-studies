package main

import "fmt"

func main() {
	// a loop can be created with the keyword `for`
	// the first parameter is the `init`, something that will execute before the loop starts (usually a variable declaration)
	// the second parameter is the `condition`, somethin that evaluates to a boolean, if it evaluates to false the loop ends
	// the third parameter is the `post`, something that will execute after each loop (usually an increment or decrement)
	for index := 0; index <= 50; index++ {
		// inside the `for` scope is the code that will run within every loop iteration
		fmt.Println(index)
	}
}
