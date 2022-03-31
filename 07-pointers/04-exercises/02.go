/*
	● create a person struct
	● create a func called “changeMe” with a *person as a parameter
		○ change the value stored at the *person address
			■ important
				● to dereference a struct, use (*value).field
					○ p1.first
					○ (*p1).first
				● “As an exception, if the type of x is a named pointer type and (*x).
						f is a valid selector expression denoting a field (but not a method), x.
						f is shorthand for (*x).f.”
						○ https://golang.org/ref/spec#Selectors
	● create a value of type person
		○ print out the value
	● in func main
		○ call “changeMe”
	● in func main
		○ print out the value
*/

package main

import "fmt"

type person struct {
	firstName, lastName string
	age                 int
}

func (p *person) changeMe() {
	p.firstName = "Dr."
	p.lastName = "No"
	p.age = 24
}

func main() {
	james := person{
		firstName: "James",
		lastName:  "Bond",
		age:       42,
	}

	fmt.Println(james)

	james.changeMe()

	fmt.Println(james)
}
