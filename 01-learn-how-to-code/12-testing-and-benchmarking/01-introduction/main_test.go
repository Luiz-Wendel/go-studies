package main

import "testing"

// to create a test function you need to name it `Test[FunctionToBeTested]` for best practice
// so, if you want to test a function called `sum` your test funtion will be called `TestSum`
// the testing function must receive a parameter of type `*testing.T` (type `T` of the `testing` package)
func TestSum(t *testing.T) {
	testCases := [][]int{{2, 3}, {4, 7}, {5, 9}}
	results := []int{5, 11, 14}

	for index, testCase := range testCases {
		x := sum(testCase...)

		if x != results[index] {
			// to raise an error you can use the `Error` method of the function parameter
			// this function is equivalent to `Log` followed by `Fail`
			t.Error("Expected:", results[index], ", Got:", x)
		}
	}
}
