// Package something is a local package with a function to sum numbers
package something

// Sum adds an unlimited number of values of type int.
func Sum(numbers ...int) int {
	result := 0

	for _, number := range numbers {
		result += number
	}

	return result
}
