package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {
	// create a `context` (ctx) with a `cancel` function
	ctx, cancel := context.WithCancel(context.Background())

	// prints out context error (should be `nil`)
	fmt.Println("error check 1:", ctx.Err())
	fmt.Println("num gortins 1:", runtime.NumGoroutine())

	// create a new goroutine
	go func() {
		number := 0

		// infinite loop
		for {
			select {
			// if context is cancelled, exit function
			case <-ctx.Done():
				return
			// else, execute some code
			default:
				number++
				time.Sleep(time.Millisecond * 200)
				fmt.Println("working", number)
			}
		}
	}()

	time.Sleep(time.Second * 2)
	fmt.Println("error check 2:", ctx.Err())
	fmt.Println("num gortins 2:", runtime.NumGoroutine())

	fmt.Println("about to cancel context")
	// cancel context
	cancel()
	fmt.Println("cancelled context")

	time.Sleep(time.Second * 2)
	fmt.Println("error check 3:", ctx.Err())
	fmt.Println("num gortins 3:", runtime.NumGoroutine())
}
