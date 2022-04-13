/*
	● write a program that:
		○ launches 10 goroutines:
			■ each goroutine adds 10 numbers to a channel;
		○ pull the numbers off the channel and print them.
*/

package main

import (
	"fmt"
	"math/rand"
	"runtime"
)

func main() {
	goroutines := 10
	numbersQuantity := 10
	c := generate(goroutines, numbersQuantity)

	for index := 0; index < goroutines*numbersQuantity; index++ {
		fmt.Println(<-c)
	}

	fmt.Println("Goroutines:", runtime.NumGoroutine())
}

func generate(goroutines, numbersQuantity int) <-chan int {
	c := make(chan int)

	for firstIndex := 0; firstIndex < goroutines; firstIndex++ {
		go func() {
			for secondIndex := 0; secondIndex < numbersQuantity; secondIndex++ {
				c <- rand.Intn(999)
			}
		}()

		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}

	return c
}
