package main

import "fmt"

var a int

type life int

var b life

func main() {
	a = 10
	b = 42
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)
	// converts `b` variable os type `life` to type `int`
	a = int(b)
	fmt.Println(a)
	fmt.Printf("%T\n", a)
}
