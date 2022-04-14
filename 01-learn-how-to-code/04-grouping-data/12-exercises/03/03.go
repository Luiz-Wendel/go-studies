/*
	Using the code from the previous example, use SLICING to create the following new slices which are then printed:
		● [42 43 44 45 46]
		● [47 48 49 50 51]
		● [44 45 46 47 48]
		● [43 44 45 46 47]
*/

package main

import "fmt"

func main() {
	slice := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	first := slice[:5]
	fmt.Println(first)

	second := slice[5:]
	fmt.Println(second)

	third := slice[2:7]
	fmt.Println(third)

	forth := slice[1:6]
	fmt.Println(forth)
}
