package main

import (
	"context"
	"fmt"
	"net/http"

	"go-chi-api/roman"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func main() {

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	CORS := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})
	r.Use(CORS.Handler)
	r.Use(middleware.SetHeader("Content-Type", "application/json"))

	r.With(myMiddleware).Get("/api/v1/test/{mydata}", MyHandler)
	r.Get("/api/v1/test/convert/", roman.RomanToInt)

	fmt.Println("Server running on port :8000")
	http.ListenAndServe(":8000", r)
}

// handler function that creates a response containing whatever is passed in {mydata}
func MyHandler(w http.ResponseWriter, r *http.Request) {
	mydata := r.Context().Value("mydata").(string)
	w.Write([]byte(fmt.Sprintf("My data is %v", mydata)))

}

// HTTP middleware setting a value on the request context
func myMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mydata := chi.URLParam(r, "mydata")
		ctx := context.WithValue(r.Context(), "mydata", mydata)
		ctxData := ctx.Value("mydata")
		fmt.Printf("My data in context is %v  \n", ctxData)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
