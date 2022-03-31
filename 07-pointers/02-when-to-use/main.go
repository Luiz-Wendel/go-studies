package main

import "fmt"

func main() {
	x := 0

	fmt.Println("x befor  ", x)

	notMutating(x)

	fmt.Println("x after z", x)

	mutating(&x)

	fmt.Println("x after y", x)
}

// this function will change the original variable value
func mutating(y *int) {
	fmt.Println("y befor  ", *y)

	*y = 42

	fmt.Println("y after  ", *y)
}

// this function won't change the original variable value
func notMutating(z int) {
	fmt.Println("z befor  ", z)

	z = 123

	fmt.Println("z after  ", z)
}
