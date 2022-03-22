package main

import "fmt"

func main() {
	for index := 33; index < 122; index++ {
		fmt.Printf("%v\t%#U\n", index, index)
	}
}
