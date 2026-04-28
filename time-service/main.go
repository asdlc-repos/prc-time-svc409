package main

import (
	"log"
	"net/http"

	"github.com/asdlc-repos/prc-time-svc409/time-service/handlers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/time", handlers.TimeHandler)
	mux.HandleFunc("/health", handlers.HealthHandler)

	addr := ":8080"
	log.Printf("time-service listening on %s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
