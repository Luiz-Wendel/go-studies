package main

import "fmt"

func main() {
	temp1 := ""

	// to invoke a function you use it's identifier followed by parentheses
	foo()

	// to call the `bar` function you must pass the correct quantity (and type) of parameters it requires inside the parentheses
	bar("James")
	name := "Moneypenny"
	bar(name)
	// the following line would throw an error `cannot use 23 (type untyped int) as type string in argument to bar`
	// bar(23)

	// to get the returned value of a function we can create a new variable or change the value of an already declared one
	temp1 = woo("James")
	fmt.Println(temp1)
	temp2 := woo(name)
	fmt.Println(temp2)

	// to get multiple values of a function we must set multiple variables
	phrase, ok := greeting("James", "Bond")
	fmt.Println(phrase)
	fmt.Println(ok)
}

// the following line declares a function with the keyword `func` and identifier `foo`
func foo() {
	fmt.Println("Hello!")
}

// to declare a function with parameters you must declare them inside the parentheses
// each parameter must have an `identifier` followed by a it's `type`
// if more than one parameter is required you must use a comma to separate them
func bar(s string) {
	fmt.Println("Hello,", s)
}

// to declare a function that returns some value to it's caller you must put the returned value `type` after the function parentheses
// and also use the `return` keyword inside the block scope followed with what you want to return
func woo(s string) string {
	return fmt.Sprint("Woo, ", s)
}

// to declare a function that returns more than one value you must open a new set of parentheses
// inside them you specify the `types`, in order, separeted by a comma
func greeting(firstName, lastName string) (string, bool) {
	phrase := fmt.Sprint(firstName, " ", lastName, `, says "Hello!"`)

	return phrase, true
}
