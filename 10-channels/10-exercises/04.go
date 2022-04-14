/*
	● Starting with [this code](https://play.golang.org/p/MvL6uamrJP), pull the values off the channel using a select statement.
*/

package main

import (
	"fmt"
)

func main() {
	q := make(chan int)
	c := gen(q)

	receive(c, q)

	fmt.Println("about to exit")
}

func gen(q chan<- int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}

		q <- 0
	}()

	return c
}

func receive(c, q <-chan int) {
	for {
		select {
		case value := <-c:
			fmt.Println(value)
		case value := <-q:
			fmt.Println("quit:", value)
			return
		}
	}
}
