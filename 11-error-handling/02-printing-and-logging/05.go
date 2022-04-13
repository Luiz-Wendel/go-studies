package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	defer foo()

	_, error := os.Open("no-file.txt")

	if error != nil {
		// equivalent as `log.Println()` followed by a call to `panic()`
		log.Panicln(error)
	}
}

func foo() {
	fmt.Println("when panic() is called, deferred functions run")
}
