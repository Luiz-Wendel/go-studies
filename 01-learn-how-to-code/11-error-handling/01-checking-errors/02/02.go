package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// function `Open` from package `os` recover a file with name as in parameter
	// returns the `file` and an `error`
	// to check an error you can pass a file name that does not exist in the directory
	file, error := os.Open("names.txt")

	if error != nil {
		fmt.Println(error)

		return
	}

	defer file.Close()

	// reads from a `Reader` and returns a `slice of bytes` and an error
	bs, error := ioutil.ReadAll(file)

	if error != nil {
		fmt.Println(error)
	}

	fmt.Println(string(bs))
}
