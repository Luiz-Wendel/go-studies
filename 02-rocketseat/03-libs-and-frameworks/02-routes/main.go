package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Timeout(time.Minute))

	// Route with parameters (with regex validation)
	// With middleware
	r.With(middleware.RealIP).
		Get("/users/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
			id := chi.URLParam(r, "id")

			fmt.Fprintln(w, id)
		})

	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}
}
