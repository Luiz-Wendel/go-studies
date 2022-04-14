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
	const goroutines = 10

	wg.Add(goroutines)

	// this will launch only 10 goroutines, instead of one for each value stored in channel `c1`
	for index := 0; index < goroutines; index++ {
		go func() {
			for v := range c1 {
				func(v2 int) {
					c2 <- timeConsumingWork(v2)
				}(v)
			}

			wg.Done()
		}()
	}

	wg.Wait()

	close(c2)
}

func timeConsumingWork(number int) int {
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(500)))

	return number + rand.Intn(1000)
}
