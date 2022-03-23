/*
	Using the code from the previous example, delete a record from your map.
	Now print the map out using the “range” loop.
*/

package main

import "fmt"

func main() {
	cast := map[string][]string{
		`bond_james`:      {`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: {`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           {`Being evil`, `Ice cream`, `Sunsets`},
	}

	cast[`sanchez_rick`] = []string{`Morty`, `Burping`, `Booze`}

	delete(cast, `no_dr`)

	for character, likes := range cast {
		fmt.Println(character, ":")
		for index, like := range likes {
			fmt.Printf("\tindex: %v\tlike: %v\n", index, like)
		}
	}
}
