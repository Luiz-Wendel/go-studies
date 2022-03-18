package main

import "fmt"

// for more info, check: https://pkg.go.dev/fmt

var y = 42

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)  // prints variable type
	fmt.Printf("%b\n", y)  // prints as binary
	fmt.Printf("%x\n", y)  // prints as hexadecimal
	fmt.Printf("%#x\n", y) // prints as hexadecimal with `0x` in front
	y = 911
	fmt.Printf("%#x\n", y)
	fmt.Printf("%#x\t%b\t%x\n", y, y, y)

	// `Sprint` means "string print", so it evaluates and returns a string with the result
	// available: `Sprint`, `Sprintln`, `Sprintf`
	// for files, check `Fprint` ("file print")
	s := fmt.Sprintf("%#x\t%b\t%x\n", y, y, y)
	fmt.Println(s)
	fmt.Printf("%v\n", y) // prints the value in a default format
}
