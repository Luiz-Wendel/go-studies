package main

import "fmt"

func main() {
	james := struct {
		firstName, lastName string
		age                 int
	}{
		firstName: "James",
		lastName:  "Bond",
		age:       32,
	}

	fmt.Println(james)
}
