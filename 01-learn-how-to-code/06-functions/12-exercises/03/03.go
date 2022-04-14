/*
	Use the “defer” keyword to show that a deferred func runs after the func containing it exits.
*/

package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("Is this the first one to run?")
	}()

	func() {
		fmt.Println("Or is this the first one to run?")
	}()
}
