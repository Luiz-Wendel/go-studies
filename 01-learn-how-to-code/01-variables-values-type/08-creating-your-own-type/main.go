package main

import "fmt"

var a int

// creates a type `life` as an integer
type life int

var b life

func main() {
	a = 10
	b = 42
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)
	// the following line would thrown an error `cannot use b (type life) as type int in assignment`
	// a = b
}
