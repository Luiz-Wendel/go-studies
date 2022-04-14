package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, error := os.Create("log.txt")

	if error != nil {
		fmt.Println(error)
	}

	defer file.Close()

	log.SetOutput(file)

	anotherFile, error := os.Open("no-file.txt")

	if error != nil {
		log.Println(error)
	}

	defer anotherFile.Close()

	fmt.Println("check the log.txt file that was created")
}
