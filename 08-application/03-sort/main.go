package main

import (
	"fmt"
	"sort"
)

func main() {
	intSlice := []int{4, 7, 3, 42, 99, 18, 16, 56, 12}
	stringSlice := []string{"James", "M", "Q", "Moneypenny", "Dr. No"}

	fmt.Println(intSlice)
	// inside package `sort` you can find the `Ints` function, which receives an `int slice` and orders it (asc)
	sort.Ints(intSlice)
	fmt.Println(intSlice)

	fmt.Println("-----")

	fmt.Println(stringSlice)
	// inside package `sort` you can find the `Strings` function, which receives a `string slice` and orders it (asc)
	sort.Strings(stringSlice)
	fmt.Println(stringSlice)
}
