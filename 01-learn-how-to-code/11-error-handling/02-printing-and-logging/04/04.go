package main

import (
	"fmt"
	"os"
)

func main() {
	defer foo()

	_, error := os.Open("no-file.txt")

	if error != nil {
		// stops normal execution of the current goroutine
		// any deferred functions run in the usual way
		panic(error)
	}
}

func foo() {
	fmt.Println("when panic() is called, deferred functions run")
}
