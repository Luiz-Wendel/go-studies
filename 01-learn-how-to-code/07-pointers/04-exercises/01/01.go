/*
	● Create a value and assign it to a variable.
	● Print the address of that value.
*/

package main

import "fmt"

func main() {
	word := "hello"

	fmt.Println(&word, word)
}
