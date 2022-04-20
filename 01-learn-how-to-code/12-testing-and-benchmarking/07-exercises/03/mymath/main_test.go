package mymath

import (
	"fmt"
	"testing"
)

func TestCenteredAvg(t *testing.T) {
	type test struct {
		data   []int
		result float64
	}

	tests := []test{
		{data: []int{1, 4, 6, 8, 100}, result: 6},
		{data: []int{0, 8, 10, 1000}, result: 9},
		{data: []int{9000, 4, 10, 8, 6, 12}, result: 9},
		{data: []int{123, 744, 140, 200}, result: 170},
	}

	for _, te := range tests {
		result := CenteredAvg(te.data)

		if result != te.result {
			t.Error("Expected: ", te.result, ", Got: ", result)
		}
	}
}

var data = []int{1, 4, 6, 8, 100}

func BenchmarkCenteredAvg(b *testing.B) {
	for index := 0; index < b.N; index++ {
		CenteredAvg(data)
	}
}

func ExampleCenteredAvg() {
	fmt.Println(CenteredAvg(data))
	// Output:
	// 6
}
