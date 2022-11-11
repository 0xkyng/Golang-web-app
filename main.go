package main

// Serving real web pages.
import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// In other for a function to respond to a request from a web browser;
// It has to handle two parameters;
// A response writer called (w http.ResponseWriter, r *http.Request)
// and a request r *http.Request

// main is the main application function
func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)


	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
