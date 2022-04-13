package main

import "fmt"

func main() {
	/*
		even := make(chan int)
		odd := make(chan int)
		// another way
		quit := make(chan bool)

		go send(even, odd, quit)

		receive(even, odd, quit)
	*/

	c := make(chan int)

	go func() {
		c <- 42

		close(c)
	}()

	// get the `value` stored in the channel and a truthy boolean indicating it has values
	value, ok := <-c

	fmt.Println(value, ok)

	// as the channel has no more values, it will return the `zero value` and a falsy boolean
	value, ok = <-c

	fmt.Println(value, ok)
}

/*
func send(even, odd chan<- int, quit chan<- bool) {
	for index := 0; index < 100; index++ {
		if index%2 == 0 {
			even <- index
		} else {
			odd <- index
		}
	}

	// closing a channel
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
*/
