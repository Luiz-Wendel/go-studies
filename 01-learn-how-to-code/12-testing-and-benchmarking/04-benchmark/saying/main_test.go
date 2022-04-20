package saying

import (
	"fmt"
	"testing"
)

func TestGreet(t *testing.T) {
	s := Greet("James")

	if s != "Hello, James" {
		t.Error("Expected: 'Hello, James', Got: ", s)
	}
}

func ExampleGreet() {
	fmt.Println(Greet("James"))
	// Output:
	// Hello, James
}

// to create a benchmark function you need to name it `Benchmark[FunctionToBeTested]` for best practice
// the benchmark function must receive a parameter of type `*testing.B` (type `B` of the `testing` package)
func BenchmarkGreet(b *testing.B) {
	// `b.N` will pick how many times the function will run
	for i := 0; i < b.N; i++ {
		Greet("James")
	}
}

/*
	to run all benchmark you need to run `go test -bench .` in the terminal
	if you want to specify the `Greet` benchmark, run `go test -bench Greet` in the terminal
*/
