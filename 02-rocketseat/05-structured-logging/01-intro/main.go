package main

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Response struct {
	Error string `json:"error,omitempty"`
	Data  any    `json:"data,omitempty"`
}

func sendJSON(w http.ResponseWriter, resp Response, status int) {
	data, err := json.Marshal(resp)
	if err != nil {
		slog.Error("error marshaling JSON", "error", err)

		sendJSON(w, Response{Error: "internal server error"}, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(status)
	if _, err := w.Write(data); err != nil {
		slog.Error("error writing JSON", "error", err)

		sendJSON(w, Response{Error: "internal server error"}, http.StatusInternalServerError)
		return
	}
}

type User struct {
	ID       int64  `json:"id,string"`
	Name     string `json:"name"`
	Role     string `json:"role"`
	Password string `json:"-"`
}

func main() {
	// SetDefault: sets the default logger for the package
	// New: creates a new logger
	// NewJSONHandler: creates a new JSON handler
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))

	// Info: logs an info level message
	// slog.Info("server started", "version", "1.0.0")

	// LogAttrs: logs an certain level message with attributes
	slog.LogAttrs(
		context.Background(),
		slog.LevelInfo,
		"server started",
		slog.String("version", "1.0.0"),
	)

	// slog.LogAttrs(
	// 	context.Background(),
	// 	slog.LevelInfo,
	// 	"http request received",
	// 	slog.String("method", http.MethodPost),
	// 	slog.String("path", "/users"),
	// 	slog.Duration("timeTaken", time.Second),
	// 	slog.Int("status", http.StatusOK),
	// )

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
			slog.Error("user not found", "id", id)

			sendJSON(w, Response{Error: "user not found"}, http.StatusNotFound)
			return
		}

		sendJSON(w, Response{Data: user}, http.StatusOK)
	}
}

func handlePostUsers(db map[int64]User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.Body = http.MaxBytesReader(w, r.Body, 1000)
		data, err := io.ReadAll(r.Body)
		if err != nil {
			var maxError *http.MaxBytesError
			if errors.As(err, &maxError) {
				slog.Error("request entity too large", "error", err)

				sendJSON(w, Response{Error: "request entity too large"}, http.StatusRequestEntityTooLarge)
				return
			}

			slog.Error("internal server error", "error", err, "body", string(data))

			sendJSON(w, Response{Error: "internal server error"}, http.StatusInternalServerError)
			return
		}

		var user User
		if err := json.Unmarshal(data, &user); err != nil {
			slog.Error("invalid body", "error", err, "body", string(data))

			sendJSON(w, Response{Error: "invalid body"}, http.StatusUnprocessableEntity)
			return
		}

		db[user.ID] = user

		sendJSON(w, Response{Data: user}, http.StatusCreated)
	}
}
