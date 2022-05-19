package main

import "fmt"

func main() {
	word := "paralelepipedo"

	fmt.Println("Word length:", len(word))

	for _, char := range word {
		fmt.Println(string(char))
	}
}
