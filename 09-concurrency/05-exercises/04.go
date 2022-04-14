/*
	Fix the race condition you created in the previous exercise by using a mutex.
		● it makes sense to remove runtime.Gosched().
*/

package main

import (
	"fmt"
	"sync"
)

var (
	wg    sync.WaitGroup
	mutex sync.Mutex
)

func main() {
	counter := 0

	gs := 50

	wg.Add(gs)

	for index := 0; index < gs; index++ {
		go incrementer(&counter)
	}

	wg.Wait()

	fmt.Println(counter)
}

func incrementer(n *int) {
	mutex.Lock()

	var number int

	number = *n

	number++

	*n = number

	mutex.Unlock()

	wg.Done()
}
