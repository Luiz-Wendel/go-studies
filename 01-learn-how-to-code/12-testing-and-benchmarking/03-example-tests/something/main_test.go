package something

import "fmt"

func ExampleSum() {
	fmt.Println(Sum(1, 2, 3))
	// Output: 6
}

/*
	Examples are compiled (and optionally executed) as part of a package’s test suite.

	As with typical tests, examples are functions that reside in a package’s _test.go files.
	Unlike normal test functions, though, example functions take no arguments and begin with the word Example instead of Test.

	As it executes the example, the testing framework captures data written to standard output and then compares the output against the example’s “Output:” comment.
	The test passes if the test’s output matches its output comment.

	Source: https://go.dev/blog/examples
*/
