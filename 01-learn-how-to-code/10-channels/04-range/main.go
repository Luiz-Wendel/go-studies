package main

import "fmt"

func main() {
	c := make(chan int)

	go foo(c)

	// this will range over the values of the channel
	for value := range c {
		fmt.Println(value)
	}

	fmt.Println("about to exit")
}

func foo(c chan<- int) {
	for i := 0; i < 100; i++ {
		c <- i
	}

	// if the channel isn't closed at the end the range will throw an error `fatal error: all goroutines are asleep - deadlock!`
	close(c)
}

/*
	? `close()` is a built-in function that is used to close a channel

	? the channel must be either bidirectional or send-only and should be executed only by the sender, never the receiver,
	? and has the effect of shutting down the channel after the last sent value is received

	? after the last value has been received from a closed channel c, any receive from c will succeed without blocking,
	? returning the zero value for the channel element.

	? source: https://www.includehelp.com/golang/close-function-with-examples.aspx
*/
