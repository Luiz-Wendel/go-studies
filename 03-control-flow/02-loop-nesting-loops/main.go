package main

import "fmt"

func main() {
	for firstIndex := 0; firstIndex <= 10; firstIndex++ {
		fmt.Printf("Outer loop: %d\n", firstIndex)

		// we can add as many nesting loops we want
		for secondIndex := 0; secondIndex < 3; secondIndex++ {
			fmt.Printf("\t\t Inner loop: %d\n", secondIndex)
		}
	}
}
