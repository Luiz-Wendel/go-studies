package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	goroutines := 100
	wg.Add(goroutines)

	// declare `counter` as `int64`
	var counter int64

	for index := 0; index < goroutines; index++ {
		go func() {
			// `AddInt64` atomically adds `delta` to `*addr` and returns the new value.
			atomic.AddInt64(&counter, 1)

			runtime.Gosched()

			// `LoadUint64` atomically loads `*addr`.
			fmt.Println("Counter\t", atomic.LoadInt64(&counter))

			wg.Done()
		}()

		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}

	wg.Wait()

	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println(counter)
}
