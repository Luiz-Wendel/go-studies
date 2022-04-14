package main

import "fmt"

func main() {
	// the `for` loop can be created with only the condition instead of `init;condition;post`
	// this is like the `while` loop in other languages
	x := 0

	for x < 10 {
		fmt.Println(x)
		x++
	}

	// an infinite loop can be created by opening the block scope right after the `for` keyword
	// to end an infinite loop we must use the `break` keyword (usually with some condition)
	y := 0

	for {
		if y > 20 {
			break
		}
		fmt.Println(y)
		y++
	}
}
