package main

import "fmt"

func main() {
	// `buffered channel` allow `n` values to stay in the channel
	// creating a `buffered channel`
	c := make(chan int, 1)

	// writes `42` to the `channel`
	c <- 42
	// since the `buffered channel` was defined with `1`, only one value can be written in the `channel`
	// the following line would throw an error `fatal error: all goroutines are asleep - deadlock!`
	// c <- 43

	// reads from the `channel`
	fmt.Println(<-c)
}

// ? for best practice, stay away from `buffered channels`, write code that interlocks writes and reads
