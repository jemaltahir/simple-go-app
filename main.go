package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello, World!")
	mux := http.NewServeMux()

	// Route for path "/" with exact match (only responds to GET requests)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprint(w, "Only GET method allowed for this endpoint")
			return
		}
		fmt.Fprint(w, "Hello, World!")
	})

	mux.HandleFunc("/hello/{name}", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Path[len("/hello/"):] // Extract substring after "/hello/"
		fmt.Fprintf(w, "Hello, %s!", name)
	})

	fmt.Println("Server listening on port 8080")
	http.ListenAndServe(":8080", mux)
}
