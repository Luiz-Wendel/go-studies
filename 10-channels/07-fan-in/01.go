package main

import (
	"fmt"
	"sync"
)

func main() {
	even := make(chan int)
	odd := make(chan int)
	fanin := make(chan int)

	go send(even, odd)

	go receive(even, odd, fanin)

	for value := range fanin {
		fmt.Println(value)
	}
}

func send(even, odd chan<- int) {
	for index := 0; index < 100; index++ {
		if index%2 == 0 {
			even <- index
		} else {
			odd <- index
		}
	}

	close(even)
	close(odd)
}

// break `receiver channels` into separate goroutines and send values into `fanin`
func receive(even, odd <-chan int, fanin chan<- int) {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for value := range even {
			fanin <- value
		}

		wg.Done()
	}()

	go func() {
		for value := range odd {
			fanin <- value
		}

		wg.Done()
	}()

	// using `WaitGroup` to ensure both functions endend executing
	wg.Wait()

	close(fanin)
}
