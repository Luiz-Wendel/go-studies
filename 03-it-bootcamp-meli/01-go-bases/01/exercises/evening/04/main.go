package main

import "fmt"

func main() {
	employees := map[string]int{"Benjamin": 20, "Manuel": 26, "Brenda": 19, "Dario": 44, "Pedro": 30}

	fmt.Println("Benjamin -", employees["Benjamin"], "year(s) old")

	var ofAge uint

	for _, employee := range employees {
		if employee > 21 {
			ofAge += 1
		}
	}

	fmt.Println("Employees of age:", ofAge)

	employees["Federico"] = 25

	fmt.Println(employees)

	delete(employees, "Pedro")

	fmt.Println(employees)
}
