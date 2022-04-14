package main

import "fmt"

func main() {
	c := make(chan int, 2)
	// receive-only channel
	cr := make(<-chan int)
	// send-only channel
	cs := make(chan<- int)

	c <- 42
	c <- 43
	// the following line would throw an error `invalid operation: cr <- 42 (send to receive-only type <-chan int)`
	// cr <- 42

	fmt.Println(<-c)
	fmt.Println(<-c)
	// the following line would throw an error `invalid operation: <-cs (receive from send-only type chan<- int)`
	// fmt.Println(<-cs)
	fmt.Println("-----")
	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", cr)
	fmt.Printf("%T\n", cs)
}
