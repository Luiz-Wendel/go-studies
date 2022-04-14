package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := fanIn(boring("Joe"), boring("Ann"))

	for index := 0; index < 10; index++ {
		fmt.Println(<-c)
	}

	fmt.Println("You're both boring; I'm leaving.")
}

func boring(msg string) <-chan string {
	c := make(chan string)

	go func() {
		for index := 0; ; index++ {
			c <- fmt.Sprintf("%s %d", msg, index)
			time.Sleep(time.Duration(time.Duration(rand.Intn(1e3)) * time.Millisecond))
		}
	}()

	return c
}

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)

	go func() {
		for {
			c <- <-input1
		}
	}()

	go func() {
		for {
			c <- <-input2
		}
	}()

	return c
}

/*
	source: Rob Pike
*/
