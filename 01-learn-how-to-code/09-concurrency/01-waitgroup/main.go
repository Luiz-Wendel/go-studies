package main

import (
	"fmt"
	"runtime"
	"sync"
)

// WaitGroup variable
var wg sync.WaitGroup

func main() {
	// check OS
	fmt.Println("OS\t", runtime.GOOS)
	// check architecture
	fmt.Println("ARCH\t", runtime.GOARCH)
	// check CPU quantity
	fmt.Println("CPUs\t", runtime.NumCPU())
	// check goroutines quantity
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	// tells WaitGroup to add `n` things to wait
	// in this case only one thing (so only one `done` will be needed)
	wg.Add(1)

	// to start a new goroutine you can use the `go` keyword
	go foo()
	bar()

	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	// tells WaitGroup to wait until everything is done executing
	wg.Wait()
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo:", i)
	}

	// when the function is done executing it tells the WaitGroup
	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar:", i)
	}
}
