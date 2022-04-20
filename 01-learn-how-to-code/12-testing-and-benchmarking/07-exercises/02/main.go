/*
	Start with [this code](https://github.com/GoesToEleven/go-programming/tree/master/code_samples/010-ninja-level-thirteen/02/01-code-starting).
	Get the code ready to BET on (add benchmarks, examples, tests) however do not write an example for the func that returns a map;
	and only write a test for it as an extra challenge.
	Add documentation to the code.

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

	"github.com/Luiz-Wendel/go-studies/01-learn-how-to-code/12-testing-and-benchmarking/07-exercises/02/quote"
	"github.com/Luiz-Wendel/go-studies/01-learn-how-to-code/12-testing-and-benchmarking/07-exercises/02/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}
