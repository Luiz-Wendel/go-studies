package main

import "fmt"

// practical application of bit shifting with iota
const (
	_  = iota
	kb = 1 << (10 * iota)
	mb = 1 << (10 * iota)
	gb = 1 << (10 * iota)
)

func main() {
	x := 1
	fmt.Printf("%d\t\t%b\n", x, x)

	// the `<<` (left shift) operator adds `n` bits `0` at the end of the binary value
	y := x << 1
	fmt.Printf("%d\t\t%b\n", y, y)

	// the `>>` (right shift) operator removes `n` bits (either `0` or `1`) at the end of the binary value
	z := y >> 1
	fmt.Printf("%d\t\t%b\n", z, z)

	n := y >> 2
	fmt.Printf("%d\t\t%b\n", n, n)

	m := y << 2
	fmt.Printf("%d\t\t%b\n", m, m)

	a := 42
	fmt.Printf("%d\t\t%b\n", a, a)

	b := a << 1
	fmt.Printf("%d\t\t%b\n", b, b)

	c := a >> 2
	fmt.Printf("%d\t\t%b\n", c, c)

	fmt.Println("kb size:")
	fmt.Printf("%d\t\t%b\n", kb, kb)
	fmt.Println("mb size:")
	fmt.Printf("%d\t\t%b\n", mb, mb)
	fmt.Println("gb size:")
	fmt.Printf("%d\t%b\n", gb, gb)
}
