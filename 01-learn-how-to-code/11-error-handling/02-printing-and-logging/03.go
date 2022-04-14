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
		// equivalent as `log.Println()` followed by a call to `os.Exit(1)`
		log.Fatalln(error)
	}
}

func foo() {
	fmt.Println("when os.Exit() is called, deferred functions don't run")
}
