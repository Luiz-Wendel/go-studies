package main

import "fmt"

func main() {
	jb := []string{"James", "Bond", "Chocolate", "Martini"}
	fmt.Println(jb)

	mp := []string{"Miss", "Moneypenny", "Strawberry", "Hazelnut"}
	fmt.Println(mp)

	// to create a multi-dimensional slice an extra pair of brackets must be placed
	// the following line creates a slice of string slice's
	xp := [][]string{jb, mp}
	fmt.Println(xp)
}
