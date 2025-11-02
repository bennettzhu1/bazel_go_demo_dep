package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// Get configuration from environment variables
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	appName := os.Getenv("APP_NAME")
	if appName == "" {
		appName = "hello-k8s"
	}

	// Simple HTTP handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		message := fmt.Sprintf("Hello from %s! 🚀\n", appName)
		fmt.Fprintf(w, message)
		log.Printf("Served request from %s", r.RemoteAddr)
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "OK\n")
	})

	log.Printf("Starting %s on port %s", appName, port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
