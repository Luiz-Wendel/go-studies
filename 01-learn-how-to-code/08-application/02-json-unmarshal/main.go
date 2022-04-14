package main

import (
	"encoding/json"
	"fmt"
)

// in the `struct` you can define a tag `jsonString` to map the fields, allowing the imported JSON to have different key names from the struct
// source: https://pkg.go.dev/encoding/json#Marshal
type person struct {
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Age       int    `json:"Age"`
}

func main() {
	jsonString := `[{"FirstName":"James","LastName":"Bond","Age":42},{"FirstName":"Miss","LastName":"Moneypenny","Age":37}]`
	// casting the string as a `byte slice`
	bs := []byte(jsonString)

	fmt.Println(jsonString)
	fmt.Println(bs)

	var people []person

	// `Unmarshal` gets a `byte slice` and the `address` where it should store the data, and returns an error (if any)
	error := json.Unmarshal(bs, &people)

	if error != nil {
		fmt.Println("Error:", error)
	}

	// prints the data
	fmt.Println(people)
}
