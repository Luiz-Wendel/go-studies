/*
	● in addition to the main goroutine, launch two additional goroutines;
		○ each additional goroutine should print something out.
	● use waitgroups to make sure each goroutine finishes before your program exists.
*/

package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	wg.Add(2)

	go func() {
		fmt.Println("First additional goroutine")

		wg.Done()
	}()

	go func() {
		fmt.Println("Second additional goroutine")

		wg.Done()
	}()

	fmt.Println("Goroutines:", runtime.NumGoroutine())

	wg.Wait()
}
