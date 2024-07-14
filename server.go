package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/":
			fmt.Fprintf(w, "This is the main page")
		case "/redirect301":
			http.Redirect(w, r, "/redirected", http.StatusMovedPermanently)
		case "/redirect302":
			http.Redirect(w, r, "/redirected", http.StatusFound)
		case "/redirected":
			fmt.Fprintf(w, "This is the redirected page")
		default:
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "Page not found")
		}
	})

	fmt.Println("Server is running on http://localhost:3000/")
	http.ListenAndServe(":3000", nil)
}
