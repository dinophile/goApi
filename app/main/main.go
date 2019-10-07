package main

import (
	"net/http"

	"github.com/dinophile/routes"
)

// Run with
//		go run .
// Send request with:
//		curl -F 'file=@/path/matrix.csv' "localhost:8080/echo"


// main - route handlers for API
func main() {
	http.HandleFunc("/echo", routes.Echo)
	http.HandleFunc("/invert", routes.Invert)
	http.HandleFunc("/flatten", routes.Flatten)
	http.HandleFunc("/sum", routes.Sum)
	http.HandleFunc("/multiply", routes.Multiply)
	http.HandleFunc("/health", routes.Health)
	http.ListenAndServe(":8080", nil)
}
