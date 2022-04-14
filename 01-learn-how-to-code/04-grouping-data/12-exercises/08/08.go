/*
	Create a map with a key of TYPE string which is a person’s “last_first” name, and a value of TYPE []string which stores their favorite things.
	Store three records in your map.
	Print out all of the values, along with their index position in the slice.
		`bond_james`, `Shaken, not stirred`, `Martinis`, `Women`
		`moneypenny_miss`, `James Bond`, `Literature`, `Computer Science`
		`no_dr`, `Being evil`, `Ice cream`, `Sunsets`
*/

package main

import "fmt"

func main() {
	cast := map[string][]string{
		`bond_james`:      {`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: {`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           {`Being evil`, `Ice cream`, `Sunsets`},
	}

	for character, likes := range cast {
		fmt.Println(character, ":")
		for index, like := range likes {
			fmt.Printf("\tindex: %v\tlike: %v\n", index, like)
		}
	}
}
