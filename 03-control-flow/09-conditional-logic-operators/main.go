package main

import "fmt"

func main() {
	// the and (`&&`) logic operator only evaluates to `true` when both sides are `true`
	fmt.Printf("true && true\t=> %v\n", true && true)
	fmt.Printf("true && false\t=> %v\n", true && false)
	fmt.Printf("false && true\t=> %v\n", false && true)
	fmt.Printf("false && false\t=> %v\n\n", false && false)

	// the or (`||`) logic operator only evaluates to `false` when both sides are `false`
	fmt.Printf("true || true\t=> %v\n", true || true)
	fmt.Printf("true || false\t=> %v\n", true || false)
	fmt.Printf("false || true\t=> %v\n", false || true)
	fmt.Printf("false || false\t=> %v\n\n", false || false)

	// the not (`!`) logic operator will toggle the boolean state
	fmt.Printf("!true \t=> %v\n", !true)
	fmt.Printf("!false \t=> %v\n", !false)
}
