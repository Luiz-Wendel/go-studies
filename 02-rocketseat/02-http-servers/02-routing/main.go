package main

import (
	"fmt"
	"net/http"
	"time"
)

// Middleware to log the requests
func Log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		begin := time.Now()

		next.ServeHTTP(w, r)

		fmt.Println(r.Method, r.URL.String(), "-", time.Since(begin))
	})
}

func main() {
	// Create a new mux (multi pattern handler)
	mux := http.NewServeMux()

	// Register a handler function for the given pattern
	// METHOD /PATTERN[.../{PATHVALUES}] ===> from version 1.22+
	mux.HandleFunc("POST /api/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		// Get method
		method := r.Method

		// Get the path values
		id := r.PathValue("id")

		fmt.Fprintln(w, "method:", method)
		fmt.Fprintln(w, "user id:", id)
	})

	// Create a new server
	srv := &http.Server{
		Addr:                         ":8080",
		Handler:                      Log(mux), // Add middleware
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
