package main

import "fmt"

func main() {
	c := make(chan int)

	// send
	go foo(c)

	// by not setting the `bar` function to another goroutine we block the execution until the `foo` goroutine runs
	// receive
	bar(c)

	fmt.Println("about to exit")
}

// send
func foo(c chan<- int) {
	c <- 42
}

// receive
func bar(c <-chan int) {
	fmt.Println(<-c)
}
