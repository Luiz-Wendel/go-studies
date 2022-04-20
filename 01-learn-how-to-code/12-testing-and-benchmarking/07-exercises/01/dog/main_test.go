package dog

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {
	type test struct {
		data   int
		result int
	}

	tests := []test{
		{data: 5, result: 35},
		{data: 2, result: 14},
		{data: 7, result: 49},
	}

	for _, te := range tests {
		result := Years(te.data)

		if result != te.result {
			t.Error("Expected: ", te.result, ", Got: ", result)
		}
	}
}

func TestYearsTwo(t *testing.T) {
	type test struct {
		data   int
		result int
	}

	tests := []test{
		{data: 5, result: 35},
		{data: 2, result: 14},
		{data: 7, result: 49},
	}

	for _, te := range tests {
		result := YearsTwo(te.data)

		if result != te.result {
			t.Error("Expected: ", te.result, ", Got: ", result)
		}
	}
}

const age = 7

func BenchmarkYears(b *testing.B) {
	for index := 0; index < b.N; index++ {
		Years(age)
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	for index := 0; index < b.N; index++ {
		YearsTwo(age)
	}
}

func ExampleYears() {
	fmt.Println(Years(age))
	// Output:
	// 49
}

func ExampleYearsTwo() {
	fmt.Println(YearsTwo(age))
	// Output:
	// 49
}
