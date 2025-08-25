package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	// get("https://google.com")

	// post("https://google.com")
}

func readBody(resp *http.Response) {
	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}

func get(url string) {
	// Get request to an url
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	// Close the response body for resources management
	defer resp.Body.Close()

	readBody(resp)
}

func post(url string) {
	// Post request to an url
	resp, err := http.Post(url, "application/json", nil)
	if err != nil {
		panic(err)
	}
	// Close the response body for resources management
	defer resp.Body.Close()

	readBody(resp)
}
