/*
	Building on the previous hands-on exercise, create a program that uses “else if” and “else”.
*/

package main

import "fmt"

func main() {
	if 42 < 0 {
		fmt.Println("this is wrong")
	} else if 42 != 42 {
		fmt.Println("this is wrong too")
	} else {
		fmt.Println("42 is the answer")
	}
}
