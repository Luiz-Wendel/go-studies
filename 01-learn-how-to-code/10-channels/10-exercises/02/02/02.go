/*
	● Get these codes running:
		○ https://play.golang.org/p/oB-p3KMiH6
		* ○ https://play.golang.org/p/_DBRueImEq
*/

package main

import (
	"fmt"
)

func main() {
	cr := make(chan int)

	go func() {
		cr <- 42
	}()
	fmt.Println(<-cr)

	fmt.Printf("------\n")
	fmt.Printf("cr\t%T\n", cr)
}
