/*
	Create TYPED and UNTYPED constants.
	Print the values of the constants.
*/

package main

import "fmt"

func main() {
	const (
		untyped        = 42
		typed   string = "42"
	)

	fmt.Println(untyped, typed)
}
