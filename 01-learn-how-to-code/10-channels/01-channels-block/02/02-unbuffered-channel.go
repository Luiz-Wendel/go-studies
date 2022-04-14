package main

import "fmt"

func main() {
	// creating a `integer channel`
	c := make(chan int)

	// create a new goroutine
	go func() {
		// writes `42` to the `channel`
		c <- 42
	}()

	// reads from the `channel`
	fmt.Println(<-c)
}
