package main

import "fmt"

// A `struct` is an composite data type (or aggregate data type or complex data type)
// `struct` allows us to create a data structure with values of different type
// to create a new `struct` we define a new type as a `struct` and put any other info inside the curly braces
type person struct {
	firstName string
	lastName  string
	age       int
}

func main() {
	james := person{
		firstName: "James",
		lastName:  "Bond",
		age:       32,
	}

	miss := person{
		firstName: "Miss",
		lastName:  "Moneypenny",
		age:       27,
	}

	fmt.Println(james)
	fmt.Println(miss)

	fmt.Println(james.firstName, james.lastName, james.age)
	fmt.Println(miss.firstName, miss.lastName, miss.age)
}
