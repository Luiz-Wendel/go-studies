/*
	● Using goroutines, create an incrementer program;
		○ have a variable to hold the incrementer value;
		○ launch a bunch of goroutines.
			■ each goroutine should.
				● read the incrementer value;
					○ store it in a new variable.
				● yield the processor with runtime.Gosched();
				● increment the new variable;
				● write the value in the new variable back to the incrementer variable.
	● use waitgroups to wait for all of your goroutines to finish;
	● the above will create a race condition;
	● Prove that it is a race condition by using the -race flag;
*/

package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

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
	var number int

	number = *n

	runtime.Gosched()

	number++

	*n = number

	wg.Done()
}

/*
	running the file with `go run -race` returned:
		Found 1 data race(s)
		exit status 66
*/
