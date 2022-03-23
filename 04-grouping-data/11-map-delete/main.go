package main

import "fmt"

func main() {
	m := map[string]int{
		"James":      32,
		"Moneypenny": 27,
	}

	fmt.Println(m)

	// the `delete` built-in function takes the `map` and the `key` to be deleted
	delete(m, "James")
	fmt.Println(m)

	// this function does not throw an error if the `key` does not exist in the `map`
	delete(m, "James")
	fmt.Println(m)

	// to assure that it only deletes a `key` that exists the `comma ok` idiom must be used
	if _, ok := m["James"]; ok {
		fmt.Println("Key found, now deleting...")
		delete(m, "James")
	}

	if _, ok := m["Moneypenny"]; ok {
		fmt.Println("Key found, now deleting...")
		delete(m, "Moneypenny")
	}
}
