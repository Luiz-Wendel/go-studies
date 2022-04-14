package main

import (
	"fmt"
	"runtime"
)

var (
	c int8  // -128 -> 127
	d uint8 // 0 -> 255
)

func main() {
	a := 42
	b := 42.4242

	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)

	// the following line would throw an error `constant 42.4242 truncated to integer`
	// a = 42.4242
	c = -128
	// the following line would throw an error `constant -129 overflows int8`
	// c = -129
	// the following line would throw an error `constant 128 overflows int8`
	// c = 128

	d = 255
	// the following line would throw an error `constant -1 overflows uint8`
	// d = -1
	// the following line would throw an error `constant 256 overflows uint8`
	// d = 256

	// prints running OS
	fmt.Println(runtime.GOOS)
	// prints running OS architecture
	fmt.Println(runtime.GOARCH)
}

/*
	byte => unint8
	rune => int32
		- rune ~ character
		- utf-8 uses 1-4 bytes to encode characters
			- 4 bytes = 4 * 8 = 32 bits
*/
