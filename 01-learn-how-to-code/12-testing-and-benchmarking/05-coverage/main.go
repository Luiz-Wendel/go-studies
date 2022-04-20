package main

import (
	"fmt"

	"github.com/Luiz-Wendel/go-studies/01-learn-how-to-code/12-testing-and-benchmarking/05-coverage/saying"
)

func main() {
	fmt.Println(saying.Greet("James"))
}

/*
	`go test -cover` prints in the terminal the coverage:
	`ok   .../12-testing-and-benchmarking/05-coverage/saying   0.001s  coverage: 100.0% of statements`

	`go test -coverprofile [filename].out` will run the coverage and export the info to a `.out` file with name specified in the command.
	after creating the file you can display the file as html with the command `go tool cover -html=[filename].out`
*/
