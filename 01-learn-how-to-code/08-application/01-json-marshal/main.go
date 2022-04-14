package main

import (
	"encoding/json"
	"fmt"
)

// to tranform in JSON you need the attributes to be in `Pascal Case`
type person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	james := person{
		FirstName: "James",
		LastName:  "Bond",
		Age:       42,
	}

	miss := person{
		FirstName: "Miss",
		LastName:  "Moneypenny",
		Age:       37,
	}

	people := []person{james, miss}

	fmt.Println(people)

	// the `Marshal` function of `encoding/json` gets a `slice` and returns a `slice of bytes` and `error` (if any)
	bs, error := json.Marshal(people)

	// for best practice, check if there was any error
	if error != nil {
		fmt.Println("Error:", error)
	}

	// you can check the `slice of bytes`
	fmt.Println(bs)
	// to get the value as a string (JSON) you can use casting
	fmt.Println(string(bs))
}
