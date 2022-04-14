package main

import "fmt"

// ! This code will throw an error `all goroutines are asleep - deadlock!`

func main() {
	// creating a `integer channel`
	c := make(chan int)

	// writes `42` to the `channel`
	c <- 42

	// reads from the `channel`
	fmt.Println(<-c)
}

/*
	? If the capacity is zero or absent, the channel is unbuffered and communication succeeds only when both a sender and receiver are ready.
	? Otherwise, the channel is buffered and communication succeeds without blocking if the buffer is not full (sends) or not empty (receives).
	? A nil channel is never ready for communication.

	Source: https://go.dev/ref/spec#Channel_types
*/
