/*
	Write a program that prints a number in decimal, binary, and hex
*/

package main

import "fmt"

func main() {
	number := 42

	fmt.Printf("Decimal: %v\n", number)
	fmt.Printf("Binary: %b\n", number)
	fmt.Printf("Hex: %#x\n", number)
}
