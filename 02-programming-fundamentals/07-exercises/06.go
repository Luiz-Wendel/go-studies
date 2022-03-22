/*
	Using iota, create 4 constants for the NEXT 4 years.
	Print the constant values.
*/

package main

import "fmt"

func main() {
	const (
		a = 2023 + iota
		b = a + iota
		c = a + iota
		d = a + iota
	)

	fmt.Println(a, b, c, d)
}
