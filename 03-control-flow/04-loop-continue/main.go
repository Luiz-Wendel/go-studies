package main

import "fmt"

func main() {
	// to skip part of the code inside the loop scope we can use the `continue` keyword
	for index := 1; index <= 100; index++ {
		fmt.Printf("%d is ", index)

		if index%2 != 0 {
			fmt.Println("odd")
			continue
		}

		// this code will only run wen the above condition is `false`
		fmt.Println("even")
	}
}
