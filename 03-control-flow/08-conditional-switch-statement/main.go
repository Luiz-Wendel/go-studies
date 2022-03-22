package main

import "fmt"

func main() {
	// ? the keyword `switch` accepts a block scope with `case` keywords followed with conditions
	// ? if the condition evaluates to `true` it will execute the code inside this case
	// ! only the first truthy case will be executed, unless a `fallthrough` keyword is specified at the end of the case
	// ! if the `fallthrough` is executed, then the next case bellow it will execute (even if evaluated to `false`)
	// ! the `fallthrough` even executes default cases
	switch {
	case false:
		fmt.Println("this will never print")
	case 42 == 24:
		fmt.Println("this should not print")
	case 42 == 42:
		fmt.Println("this should print")
		// if the following line is not commented, the next case will also execute, even if not evaluated to `true`
		fallthrough
	case false:
		fmt.Println("this should not print")
	case "xablau" == "xablau":
		fmt.Println("will this print?")
	}

	switch {
	case false:
		fmt.Println("this will not print")
	// the following code will only execute if no case was evaluated to `true` or the case above executed and had a `fallthrough` keyword
	default:
		fmt.Println("default")
	}

	// ? switch statements can alse be used to evaluate some variable
	x := "Xablau"

	switch x {
	case "James":
		fmt.Println("James Bond")
	case "Xablau":
		fmt.Println("Xablau")
	// multiple values can be associated with the same block scope by using a comma separator
	case "John", "Jane":
		fmt.Println("Doe")
	default:
		fmt.Println("Not matched")
	}
}
