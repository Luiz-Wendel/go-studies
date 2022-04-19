package main

import "testing"

func TestSum(t *testing.T) {
	// creating a `test` struct
	type test struct {
		// defining which type of data we will test
		data []int
		// defining the results we are looking for
		answer int
	}

	// creating the tests cases (a slice of `tests`)
	tests := []test{
		{data: []int{21, 21}, answer: 42},
		{data: []int{3, 4, 5}, answer: 12},
		{data: []int{1, 1}, answer: 2},
		{data: []int{-1, 0, 1}, answer: 0},
	}

	// ranging over the tests cases
	for _, value := range tests {
		// testing the data value
		result := sum(value.data...)

		// checking if the result is equal to the expected value
		if result != value.answer {
			t.Error("Expected:", value.answer, ", Got:", result)
		}
	}
}
