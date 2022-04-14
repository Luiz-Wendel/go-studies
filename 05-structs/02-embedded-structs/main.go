package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}

// you can embed another `struct` by typing only it's type
type secretAgent struct {
	person
	ltk bool
}

// or, if you want more than one, by defining with identifiers
type doubleAgent struct {
	firstIdentity  person
	secondIdentity person
	ltk            bool
}

func main() {
	// creating from scratch
	james := secretAgent{
		person: person{
			firstName: "James",
			lastName:  "Bond",
			age:       32,
		},
		ltk: true,
	}

	miss := person{
		firstName: "Miss",
		lastName:  "Moneypenny",
		age:       27,
	}

	// creating with an instance of `person`
	missAgent := secretAgent{
		person: miss,
		ltk:    false,
	}

	samuel := person{
		firstName: "Samuel",
		lastName:  "Morland",
		age:       39,
	}

	richard := person{
		firstName: "Richard",
		lastName:  "Willis",
		age:       51,
	}

	// by omitting attributes they will receive it's zero value
	mole := doubleAgent{
		firstIdentity:  samuel,
		secondIdentity: richard,
	}

	fmt.Println(james)
	fmt.Println(missAgent)

	// the `person` attributes get promoted to the same scope as the `secretAgent`
	// but if there is any name collision you will need to specify with dot notation, like the `james` print
	fmt.Println(james.person.firstName, james.person.lastName, james.person.age, james.ltk)
	fmt.Println(missAgent.firstName, missAgent.lastName, missAgent.age, missAgent.ltk)
	fmt.Println(mole.firstIdentity.firstName, mole.firstIdentity.lastName, mole.firstIdentity.age, mole.ltk)
	fmt.Println(mole.secondIdentity.firstName, mole.secondIdentity.lastName, mole.secondIdentity.age, mole.ltk)
}
