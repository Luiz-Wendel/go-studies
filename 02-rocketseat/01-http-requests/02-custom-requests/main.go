package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	// request("https://google.com", http.MethodGet)

	// initServer()
	// requestWithTimeoutContext("http://localhost:8080", http.MethodGet)
}

func initServer() {
	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(6 * time.Second)
		fmt.Fprintln(w, "Hello World")
	}))

	go func() {
		if err := http.ListenAndServe(":8080", nil); err != nil {
			panic(err)
		}
	}()
}

func request(url string, method string) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		panic(err)
	}

	// Both will create a new header if it doesn't exist, difference:
	// Add: if header already exists, it will be added
	// Set: if header already exists, it will be replaced
	// req.Header.Add("authorization", "123")
	req.Header.Set("Accept", "application/json")

	// send request with default client
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// read response body
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))
}

func requestWithTimeoutContext(url string, method string) {
	// create context
	ctx := context.Background()
	// add timeout to context
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	// cancel context when done
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, method, url, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))
}
