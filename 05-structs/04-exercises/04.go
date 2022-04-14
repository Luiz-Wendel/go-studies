/*
	Create and use an anonymous struct.
*/

package main

import "fmt"

func main() {
	dunno := struct {
		name   string
		isTrue bool
	}{
		name:   "Xablau",
		isTrue: false,
	}

	fmt.Println(dunno)
	fmt.Printf("%T\n", dunno)
}
