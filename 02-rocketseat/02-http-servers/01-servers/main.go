package main

import (
	"fmt"
	"net/http"
	"time"
)

// We will create a server that will be exposed to the internet
// There is another aproach using reverse proxy

func main() {
	// Create a new mux (multi pattern handler)
	mux := http.NewServeMux()

	// Register a handler function for the given pattern
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "pong")
	})

	// Create a new server
	srv := &http.Server{
		Addr:                         ":8080",
		Handler:                      mux,
		DisableGeneralOptionsHandler: false,
		ReadTimeout:                  10 * time.Second,
		WriteTimeout:                 10 * time.Second,
		IdleTimeout:                  1 * time.Minute,
	}

	// Start the server
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		panic(err)
	}
}
