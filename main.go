package main

import (
	"MyPackages/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/subscription", handlers.SubscriptionHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

}
