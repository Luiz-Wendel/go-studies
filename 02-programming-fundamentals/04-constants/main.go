package main

import "fmt"

// untyped constants
/*
	const (
		a = 42
		b = 42.42
		c = "James Bond"
	)
*/

// typed constants
const (
	a int     = 42
	b float64 = 42.42
	c string  = "James Bond"
)

func main() {
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)
	fmt.Println(c)
	fmt.Printf("%T\n", c)
}
