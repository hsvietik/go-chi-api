package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/api/v1/test/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	fmt.Println("Server running on port :8000")
	http.ListenAndServe(":8000", r)
}