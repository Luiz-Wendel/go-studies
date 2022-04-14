// Entry point of program is `package main`, `func main`
package main

import "fmt"

func main() {
	// (1) sequence
	fmt.Println("Hello World!")
	foo()
	fmt.Println("Something more")

	// (2) loop/iterative
	for i := 0; i < 25; i++ {
		// (3) conditional
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

func foo() {
	fmt.Println("I'm foo.")
}

// Control flow:
// (1) sequence;
// (2) loop/iterative;
// (3) conditional.
