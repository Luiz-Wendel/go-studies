// package name
package main

// imported package
import "fmt"

func main() {
	// `package.identifier`
	// from package `fmt`, use the`Println()`
	// get the return values of `Println()` (bytes and error) and store in variables
	// `_` means that we know there's a value there but we will not use it
	bytes, _ := fmt.Println("Hello!", 12, true)
	fmt.Println(bytes)
}
