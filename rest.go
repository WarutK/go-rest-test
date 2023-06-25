package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Define the handler function for the /hello endpoint
	helloHandler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	}

	// Define the handler function for the /goodbye endpoint
	goodbyeHandler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Goodbye, World!")
	}

	// Register the handler functions with the respective endpoints
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/goodbye", goodbyeHandler)

	// Start the HTTP server and listen on port 8080
	fmt.Println("Server listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// go mod init wrt-rest-test
// go run rest.go
