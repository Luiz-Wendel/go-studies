package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// function `Create` from package `os` creates a file with name as in parameter
	// returns the `file` and an `error`
	file, error := os.Create("names.txt")

	// check if any error
	if error != nil {
		fmt.Println(error)
		return
	}

	// closes the file, renderingit unusable for I/O
	defer file.Close()

	// returns a new `Reader` reading from string parameter
	text := strings.NewReader("James Bond")

	// copies some `Reader` (second parameter) to a `Writer`
	io.Copy(file, text)
}
