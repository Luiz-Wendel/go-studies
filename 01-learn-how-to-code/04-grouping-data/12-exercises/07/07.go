/*
	Create a slice of a slice of string ([][]string).
	Store the following data in the multi-dimensional slice:
		"James", "Bond", "Shaken, not stirred"
		"Miss", "Moneypenny", "Helloooooo, James."

	Range over the records, then range over the data in each record.
*/

package main

import "fmt"

func main() {
	var md [][]string

	james := []string{"James", "Bond", "Shaken, not stirred"}
	miss := []string{"Miss", "Moneypenny", "Helloooooo, James."}

	md = append(md, james)
	md = append(md, miss)

	for firstIndex, record := range md {
		fmt.Println("Record:", firstIndex)
		for secondIndex, data := range record {
			fmt.Printf("\tIndex: %v\tValue: %v\n", secondIndex, data)
		}
	}
}
