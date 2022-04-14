/*
	Create your own type “person” which will have an underlying type of “struct” so that it can store the following data:
		● first name
		● last name
		● favorite ice cream flavors

	Create two VALUES of TYPE person.
	Print out the values, ranging over the elements in the slice which stores the favorite flavors.
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

	fmt.Println(jane.firstName, jane.lastName)
	for index, flavor := range jane.favoriteIceCreamFlavors {
		fmt.Printf("\t%v: %v\n", index+1, flavor)
	}

	fmt.Println(john.firstName, john.lastName)
	for index, flavor := range john.favoriteIceCreamFlavors {
		fmt.Printf("\t%v: %v\n", index+1, flavor)
	}
}
