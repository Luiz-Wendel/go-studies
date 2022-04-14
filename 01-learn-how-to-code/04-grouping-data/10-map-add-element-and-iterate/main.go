package main

import "fmt"

func main() {
	m := map[string]int{
		"James":      32,
		"Moneypenny": 27,
	}

	fmt.Println(m)

	// adding new `key` "Todd" to map with `value` 33
	m["Todd"] = 33

	// to iterate over the map keys and values we can use the `range` keyword
	for key, value := range m {
		fmt.Println(key, value)
	}
}
