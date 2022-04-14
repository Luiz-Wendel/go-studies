/*
	● Create a user defined struct with;
		○ the identifier “person”;
		○ the fields:
			■ firstName;
			■ lastName;
			■ age.
	● attach a method to type person with;
		○ the identifier “speak”;
		○ the method should have the person say their name and age.
	● create a value of type person;
	● call the method from the value of type person.
*/

package main

import "fmt"

type person struct {
	firstName, lastName string
	age                 int
}

func (p person) speak() {
	fmt.Println("Hi, I'm", p.firstName, p.lastName, "-", p.age, "years old.")
}

func main() {
	james := person{
		"James",
		"Bond",
		42,
	}

	james.speak()
}
