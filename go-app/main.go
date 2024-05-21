package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from GO! - test GO CI")
}

func main() {
	http.HandleFunc("/", helloWorldHandler)
	fmt.Println("Server is running on http://localhost:8080")
	err := http.ListenAndServe(":9080", nil)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
