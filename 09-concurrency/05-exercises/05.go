/*
	Fix the race condition you created in exercise #3 by using package atomic.
*/

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

func main() {
	var counter int64

	gs := 50

	wg.Add(gs)

	for index := 0; index < gs; index++ {
		go incrementer(&counter)
	}

	wg.Wait()

	fmt.Println(counter)
}

func incrementer(n *int64) {
	atomic.AddInt64(n, 1)

	wg.Done()
}

/*
	running the file with `go run -race` returned:
		Found 1 data race(s)
		exit status 66
*/
