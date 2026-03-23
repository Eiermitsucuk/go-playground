package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		chain := logMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello Client")
		fmt.Printf("Request at %v\n", time.Now())
		for k, v := range r.Header {
			fmt.Printf("%v: %v\n", k, v)
		}
		next.ServeHTTP(w, r)
	}))
	return logMiddleware(chain)
}

func main() {
	http.HandlemFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
