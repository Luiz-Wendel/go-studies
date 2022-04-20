/*
	Start with [this code](https://github.com/GoesToEleven/go-programming/tree/master/code_samples/010-ninja-level-thirteen/01/starting-code).
	Get the code ready to BET on the code (add benchmarks, examples, tests).

	Run the following in this order:
		● tests
		● benchmarks
		● coverage
		● coverage shown in web browser
		● examples shown in documentation in a web browser
*/

package main

import (
	"fmt"

	"github.com/Luiz-Wendel/go-studies/01-learn-how-to-code/12-testing-and-benchmarking/07-exercises/01/dog"
)

type canine struct {
	name string
	age  int
}

func main() {
	fido := canine{
		name: "Fido",
		age:  dog.Years(10),
	}

	fmt.Println(fido)
	fmt.Println(dog.YearsTwo(20))
}
