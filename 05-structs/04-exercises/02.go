/*
	Take the code from the previous exercise, then store the values of type person in a map with the key of last name.
	Access each value in the map.
	Print out the values, ranging over the slice.
*/

package main

import "fmt"

type person struct {
	firstName, lastName     string
	favoriteIceCreamFlavors []string
}

func main() {
	jane := person{
		firstName:               "Jane",
		lastName:                "Doe",
		favoriteIceCreamFlavors: []string{`chocolate`, `mint`},
	}

	john := person{
		firstName:               "John",
		lastName:                "Doe",
		favoriteIceCreamFlavors: []string{`chocolate`, `vanilla`, `strawberry`},
	}

	// saved the key as the `firstName` instead of `lastName` to avoid collision
	people := map[string]person{
		jane.firstName: jane,
		john.firstName: john,
	}

	for _, value := range people {
		fmt.Println(value.firstName, value.lastName)
		for index, flavor := range value.favoriteIceCreamFlavors {
			fmt.Printf("\t%v: %v\n", index+1, flavor)
		}
	}
}
