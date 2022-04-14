package main

import "fmt"

func main() {
	// functions can be assigned to a variable
	f := func() {
		fmt.Println("My first func expression")
	}

	// to call it you mus use the variable identifier with a pair of parentheses at the end
	f()
}
