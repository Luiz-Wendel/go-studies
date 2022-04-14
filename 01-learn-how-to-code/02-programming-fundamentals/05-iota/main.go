package main

import "fmt"

const (
	a = iota // creates enumerated constants
	b
	c
	// the following lines won't reset, only keep incrementing
	/*
		d = iota
		e
		f
	*/
)

// the following lines will reset
const (
	d = iota
	e
	f
)

func main() {
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Println(c)

	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
}
