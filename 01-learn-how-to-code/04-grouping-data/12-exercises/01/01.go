/*
	● Using a COMPOSITE LITERAL:
		○ create an ARRAY which holds 5 VALUES of TYPE int
		○ assign VALUES to each index position.
	● Range over the array and print the values out.
	● Using format printing
		○ print out the TYPE of the array
*/

package main

import "fmt"

func main() {
	array := [5]int{1, 2, 3, 4, 5}

	for _, number := range array {
		fmt.Println(number)
	}

	fmt.Printf("%T\n", array)
}
