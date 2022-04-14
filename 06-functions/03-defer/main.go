package main

import "fmt"

func main() {
	// deferred functions are invoked immediately before the surrounding function returns, in the reverse order they were deferred
	// deferred functions are executed after any result parameters are set by that return statement but before the function returns to its caller
	defer closeFile()
	openFile()
}

func openFile() {
	fmt.Println("File open")
}

func closeFile() {
	fmt.Println("File closed")
}
