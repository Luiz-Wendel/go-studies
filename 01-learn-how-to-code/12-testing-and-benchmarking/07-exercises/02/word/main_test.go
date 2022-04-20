package word

import (
	"fmt"
	"testing"

	"github.com/Luiz-Wendel/go-studies/01-learn-how-to-code/12-testing-and-benchmarking/07-exercises/02/quote"
)

func TestCount(t *testing.T) {
	type test struct {
		data   string
		result int
	}

	tests := []test{
		{data: "James Bond", result: 2},
		{data: "Stirred, not shaken", result: 3},
		{data: quote.SunAlso, result: 1349},
	}

	for _, te := range tests {
		result := Count(te.data)

		if result != te.result {
			t.Error("Expected: ", te.result, ", Got: ", result)
		}
	}
}

func TestUseCount(t *testing.T) {
	result := UseCount("one two two three three three")

	for key, value := range result {
		switch key {
		case "one":
			if value != 1 {
				t.Error("Expected: ", 1, ", Got: ", value)
			}
		case "two":
			if value != 2 {
				t.Error("Expected: ", 2, ", Got: ", value)
			}
		case "three":
			if value != 3 {
				t.Error("Expected: ", 3, ", Got: ", value)
			}
		}
	}
}

func BenchmarkCount(b *testing.B) {
	for index := 0; index < b.N; index++ {
		Count(quote.SunAlso)
	}
}

func BenchmarkUseCount(b *testing.B) {
	for index := 0; index < b.N; index++ {
		UseCount(quote.SunAlso)
	}
}

func ExampleCount() {
	fmt.Println(Count("James Bond"))
	// Output:
	// 2
}

func ExampleUseCount() {
	fmt.Println(UseCount("one two two three three three"))
	// Output:
	// map[one:1 three:3 two:2]
}
