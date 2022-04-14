package main

import "fmt"

type person struct {
	firstName, lastName string
}

type secretAgent struct {
	person
	ltk bool
}

func (s secretAgent) speak() {
	fmt.Println("I am", s.firstName, s.lastName, "- said the secret agent")
}

// polymorphism is when a function have diferent mechanics depending on which object (type here in Go) called the function
func (p person) speak() {
	fmt.Println("I am", p.firstName, p.lastName, "- said the person")
}

// to create an interface you need to define a new `type` as an `interface` type
// inside the block scope of an `interface` you define which methods the variable needs to have to be of the new interface type
// if the method returns something, you have to specify the return type for the method as well. ex.: speak() string
type human interface {
	speak()
}

// every variable that has the `speak` method will be of `type` human

// the following function receives a human interface as a parameter
func bar(h human) {
	// normally you switch on value of variable, but Go allows you to switch on type of variable
	switch h.(type) {
	case person:
		// you can call the interface methods directly
		h.speak()
		// but you need to use assertion to access others attributes/methods outside the interface
		fmt.Println("I'm a human person,", h.(person).firstName, h.(person).lastName)
	case secretAgent:
		h.speak()
		fmt.Println("I'm a human secret agent,", h.(secretAgent).firstName, h.(secretAgent).lastName)
	default:
		fmt.Println("I'm a human", h)
	}
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

	dr := person{
		"Dr",
		"No",
	}

	// in Go, a variable can be of more than one type
	bar(james)
	bar(miss)
	bar(dr)
}
