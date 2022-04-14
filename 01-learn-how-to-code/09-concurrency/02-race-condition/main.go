package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	goroutines := 100
	wg.Add(goroutines)

	counter := 0

	for index := 0; index < goroutines; index++ {
		// all routines will get the `counter` value and increment it
		go func() {
			fmt.Println("Goroutines:", runtime.NumGoroutine())

			x := counter

			x++

			// Gosched yields the processor, allowing other goroutines to run.
			// It does not suspend the current goroutine, so execution resumes automatically.
			runtime.Gosched()

			counter = x

			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("Goroutines:", runtime.NumGoroutine())
	// counter that should be set to `100` gives out a different number because of race condition
	fmt.Println(counter)
}

//? to check if code have any data race you can run the code with `go run -race [file_name]`
