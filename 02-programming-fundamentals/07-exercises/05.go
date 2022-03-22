/*
	Create a variable of type string using a raw string literal.
	Print it.
*/

package main

import "fmt"

func main() {
	s := `raw
"string"
literal`

	fmt.Println(s)
}
