package main

import (
	"fmt"
	"sort"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

// ! watch out for the Pascal Case
// begin creating a new type as a slice of the type which you want to sort
type ByAge []Person

// create a method `Len` that returns the length
func (a ByAge) Len() int { return len(a) }

// create a method `Swap` which swaps two values
func (a ByAge) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

// create a method `Less` that returns a boolean indicating if the first parameter is less than the second
// to compare `Age` the `<` operator was used, but any logic can be made as logn as it returns a boolean
// so depeneding on your need you will be using different operators and logic here
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

// second example
type ByLastName []Person

func (n ByLastName) Len() int           { return len(n) }
func (n ByLastName) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }
func (n ByLastName) Less(i, j int) bool { return n[i].LastName < n[j].LastName }

func main() {
	p1 := Person{
		FirstName: "James",
		LastName:  "Bond",
		Age:       32,
	}

	p2 := Person{
		FirstName: "Miss",
		LastName:  "Moneypenny",
		Age:       27,
	}

	p3 := Person{
		FirstName: "Dr.",
		LastName:  "Yes",
		Age:       64,
	}

	p4 := Person{
		FirstName: "Dr.",
		LastName:  "No",
		Age:       56,
	}

	people := []Person{p1, p2, p3, p4}

	fmt.Println(people)

	// when sorting you should cast the slice of type to the new type you created
	sort.Sort(ByAge(people))
	fmt.Println(people)

	sort.Sort(ByLastName(people))
	fmt.Println(people)
}
