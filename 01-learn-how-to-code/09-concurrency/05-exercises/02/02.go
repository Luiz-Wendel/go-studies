/*
	This exercise will reinforce our understanding of method sets:
		● create a type person struct;
		● attach a method speak to type person using a pointer receiver;
			○ *person.
		● create a type human interface;
			○ to implicitly implement the interface, a human must have the speak method.
		● create func “saySomething”;
			○ have it take in a human as a parameter;
			○ have it call the speak method.
		● show the following in your code;
			○ you CAN pass a value of type *person into saySomething;
			○ you CANNOT pass a value of type person into saySomething.
*/

package main

import "fmt"

type person struct {
	firstName, lastName string
}

func (p *person) speak() {
	fmt.Println("I'm,", p.firstName, p.lastName)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	james := person{
		firstName: "James",
		lastName:  "Bond",
	}

	// the following line would throw an error `cannot use james (type person) as type human in argument to saySomething`
	// saySomething(james)

	saySomething(&james)
}
