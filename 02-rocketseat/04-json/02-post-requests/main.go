package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type User struct {
	ID       int64  `json:"id,string"`
	Name     string `json:"name"`
	Role     string `json:"role"`
	Password string `json:"-"` // ignore field on JSON Marshal/Unmarshal
}

func main() {
	r := chi.NewMux()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	db := map[int64]User{
		1: {
			ID:       1,
			Name:     "admin",
			Role:     "admin",
			Password: "admin",
		},
	}

	r.Group(func(r chi.Router) {
		r.Use(jsonMiddleware)

		r.Get("/users/{id:[0-9]+}", handleGetUsers(db))

		r.Post("/users", handlePostUsers(db))
	})

	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}
}

func jsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		next.ServeHTTP(w, r)
	})
}

func handleGetUsers(db map[int64]User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, _ := strconv.ParseInt(idStr, 10, 64)

		user, ok := db[id]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`{"msg": "user not found"}`))
			return
		}

		data, err := json.Marshal(user)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"msg": "internal server error"}`))
			return
		}

		w.Write(data)
	}
}

func handlePostUsers(db map[int64]User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.Body = http.MaxBytesReader(w, r.Body, 1000)
		data, err := io.ReadAll(r.Body)
		if err != nil {
			var maxError *http.MaxBytesError
			if errors.As(err, &maxError) {
				w.WriteHeader(http.StatusRequestEntityTooLarge)
				w.Write([]byte(`{"msg": "request entity too large"}`))
				return
			}

			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"msg": "internal server error"}`))
			return
		}

		var user User
		if err := json.Unmarshal(data, &user); err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			w.Write([]byte(`{"msg": "invalid body"}`))
			return
		}

		db[user.ID] = user

		w.WriteHeader(http.StatusCreated)
		w.Write(data)
	}
}
