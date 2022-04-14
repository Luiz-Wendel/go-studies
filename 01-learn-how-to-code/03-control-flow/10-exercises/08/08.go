/*
	Create a program that uses a switch statement with no switch expression specified.
*/

package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("this won't print")
	case true:
		fmt.Println("this will print")
		fallthrough
	default:
		fmt.Println("this will print because of the fallthrough")
	}
}
