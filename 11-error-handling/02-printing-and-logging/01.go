package main

import (
	"log"
	"os"
)

func main() {
	_, error := os.Open("no-file.txt")

	if error != nil {
		// fmt.Println("error:", error)

		// same as `fmt.Println` but with datetime preceding output
		log.Println("error:", error)
	}
}
