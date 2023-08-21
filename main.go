package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.With(MyMiddleware).Get("/api/v1/test/{mydata}", MyHandler)

	fmt.Println("Server running on port :8000")
	http.ListenAndServe(":8000", r)
}

//handler function that creates a response containing whatever is passed in {mydata} 
func MyHandler(w http.ResponseWriter, r *http.Request){
	mydata := chi.URLParam(r, "mydata") 
	w.Write([]byte(fmt.Sprintf("My data is %v", mydata)))
}


// HTTP middleware setting a value on the request context
func MyMiddleware( next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	  mydata := chi.URLParam(r, "mydata")
	  ctx := context.WithValue(r.Context(), "mydata", mydata)
	  ctxData:=ctx.Value("mydata")
	  fmt.Printf("My data in context is %v  \n", ctxData)
  	  next.ServeHTTP(w, r.WithContext(ctx))
	})
}