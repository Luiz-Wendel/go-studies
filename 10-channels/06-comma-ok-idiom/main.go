package main

import "fmt"

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)

	go send(even, odd, quit)

	receive(even, odd, quit)
}

func send(even, odd chan<- int, quit chan<- bool) {
	for index := 0; index < 100; index++ {
		if index%2 == 0 {
			even <- index
		} else {
			odd <- index
		}
	}

	close(quit)
}

func receive(even, odd <-chan int, quit <-chan bool) {
	for {
		select {
		case value := <-even:
			fmt.Println("Even:\t", value)
		case value := <-odd:
			fmt.Println("Odd:\t", value)
		case value := <-quit:
			fmt.Println("Quit:\t", value)
			return
		}
	}
}
