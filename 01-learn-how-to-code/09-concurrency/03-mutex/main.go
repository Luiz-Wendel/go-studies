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

	// create the `mutex` variable
	var mu sync.Mutex

	for index := 0; index < goroutines; index++ {
		go func() {
			// `Lock` locks the `mutex`.
			// If the lock is already in use, the calling goroutine blocks until the mutex is available.
			mu.Lock()

			x := counter

			runtime.Gosched()

			x++

			counter = x

			// `Unlock` unlocks the `mutex`.
			// It is a run-time error if the `mutex` is not locked on entry to `Unlock`.
			mu.Unlock()

			wg.Done()
		}()

		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}

	wg.Wait()

	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println(counter)
}

/*
	A RWMutex is a reader/writer mutual exclusion lock.
	The lock can be held by an arbitrary number of readers or a single writer.

	If a goroutine holds a RWMutex for reading and another goroutine might call Lock, no goroutine should expect to be able to acquire a read lock until the initial read lock is released.
	In particular, this prohibits recursive read locking.
	This is to ensure that the lock eventually becomes available; a blocked Lock call excludes new readers from acquiring the lock.

	Source: https://pkg.go.dev/sync#RWMutex
*/
