package main

import "fmt"

type person struct {
	firstName, lastName string
}

type secretAgent struct {
	person
	ltk bool
}

/*
	a method is a function with a receiver
	a method declaration binds an identifier to a method, and associates the method with the receiver's base type
	The receiver is specified via an extra parameter section preceding the method name
	That parameter section must declare a single non-variadic parameter, the receiver
	The method is said to be bound to its receiver base type and the method name is visible only within selectors for type T or *T

	Source: https://go.dev/ref/spec#Method_declarations
*/
// func (r receiver) identifier(parameters) (return(s)) { ... }
func (s secretAgent) speak() {
	fmt.Println("I am", s.firstName, s.lastName)
}

func main() {
	james := secretAgent{
		person: person{
			"James",
			"Bond",
		},
		ltk: true,
	}

	miss := secretAgent{
		person: person{
			"Miss",
			"Moneypenny",
		},
		ltk: true,
	}

	fmt.Println(james)
	// to call a method, a variable of `type` defined as the method `receiver` must be used followed by dot notation and the method name with parentheses
	james.speak()
	miss.speak()
}
