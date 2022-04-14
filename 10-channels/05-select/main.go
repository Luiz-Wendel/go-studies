package main

import "fmt"

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	go create(even, odd, quit)

	// controller for infinite loop
	// exit := false

	// infinite loop
	for {
		// `select` is the same as `switch` but for `channels`
		select {
		case value := <-even:
			fmt.Println("Even:\t", value)
		case value := <-odd:
			fmt.Println("Odd:\t", value)
		// quit case, to stop the infinite loop
		case value := <-quit:
			fmt.Println("Quit:\t", value)
			// `return` keyword will break the `for` loop by returning to the `main` function
			return
			// exit = true
		}

		// if exit {
		// 	break
		// }
	}

	// any code past here will run only with the controller variable and break condition inside `for` loop
	// fmt.Println("about to exit")
}

func create(even, odd, quit chan<- int) {
	for index := 0; index < 100; index++ {
		if index%2 == 0 {
			even <- index
		} else {
			odd <- index
		}
	}

	quit <- 0
}

/*
	? A `select` will choose multiple valid options at random,
	? while a `switch` will go in sequence (and would require a fallthrough to match multiple)

	? source: https://stackoverflow.com/questions/38821491/what-is-the-difference-between-switch-and-select-in-go
*/
