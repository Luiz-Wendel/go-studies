package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go populate(c1)

	go fanOutIn(c1, c2)

	for value := range c2 {
		fmt.Println(value)
	}

	fmt.Println("about to exit")
}

func populate(c chan int) {
	for i := 0; i < 100; i++ {
		c <- i
	}

	close(c)
}

func fanOutIn(c1, c2 chan int) {
	var wg sync.WaitGroup

	// `fan out`, make a goroutine for every value in channel `c1`
	for value := range c1 {
		wg.Add(1)

		go func(anotherValue int) {
			// `fan in`, get the result value and store in channel `c2`
			c2 <- timeConsumingWork(anotherValue)

			wg.Done()
		}(value)
	}

	wg.Wait()

	close(c2)
}

func timeConsumingWork(number int) int {
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(500)))

	return number + rand.Intn(1000)
}
