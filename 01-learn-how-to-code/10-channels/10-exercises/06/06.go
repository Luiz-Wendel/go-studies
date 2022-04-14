/*
	● write a program that:
		○ puts 100 numbers to a channel;
		○ pull the numbers off the channel and print them.
*/

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	c := make(chan int)

	go generate(c)

	for value := range c {
		fmt.Println(value)
	}
}

func generate(c chan<- int) {
	for index := 0; index < 100; index++ {
		c <- rand.Intn(999)
	}

	close(c)
}
