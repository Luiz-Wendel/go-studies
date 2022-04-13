/*
	● Starting with [this code](https://play.golang.org/p/sfyu4Is3FG), pull the values off the channel using a for range loop.
*/

package main

import (
	"fmt"
)

func main() {
	c := gen()
	receive(c)

	fmt.Println("about to exit")
}

func gen() <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}

		close(c)
	}()

	return c
}

func receive(c <-chan int) {
	for value := range c {
		fmt.Println(value)
	}
}
