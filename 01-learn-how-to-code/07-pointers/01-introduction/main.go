package main

import "fmt"

func main() {
	a := 42

	// accessing the value `42`
	fmt.Println(a)
	// accessing the memory address `0x[...]`
	fmt.Println(&a)
	// accessing the type of the variable `int` (integer)
	fmt.Printf("%T\n", a)
	// accessing the type of the address `*int` (pointer to integer)
	fmt.Printf("%T\n", &a)

	// the following line would throw an error `cannot use &a (type *int) as type int in assignment`
	// var b int = &a

	var b *int = &a
	fmt.Println(b)

	// accessing the value stored at an address (dereferencing an address) `42`
	fmt.Println(*b)

	// changing the value stored at an address
	*b = 10
	// since `a` and `b` are referencing the same address, both values change
	fmt.Println(a)
	fmt.Println(*b)
}
