/*
  Closure is when we have “enclosed” the scope of a variable in some code block.
  For this hands-on exercise, create a func which “encloses” the scope of a variable.
*/

package main

import "fmt"

func main() {
	fmt.Println("Closure example")

	{
		number := 42

		fmt.Println(number)
	}

	number := 10

	fmt.Println(number)
}
