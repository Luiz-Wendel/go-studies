package main

import "fmt"

func main() {
	// composite literal to create a map
	// the map type is defined by the `map` keyword, the `key type` inside the brackets and `value type` after the brackets
	m := map[string]int{
		"James":      32,
		"Moneypenny": 27,
	}

	fmt.Println(m)

	// a `value` can be accessed through it's `key`, like in arrays and slices
	fmt.Println(m["James"])

	// if an invalid `key` is accessed, the `zero value` of the `value type` will be returned
	fmt.Println(m["Bond"])

	// to check if a key exists in the `map` we can use the `comma ok` idiom
	value, ok := m["Bond"]
	fmt.Println(value)
	fmt.Println(ok)

	// a common practice:
	if v, ok := m["Bond"]; ok {
		fmt.Println("The key exists, value:", v)
	}

	if v, ok := m["Moneypenny"]; ok {
		fmt.Println("The key exists, value:", v)
	}
}
