package main

import (
	"geteway/internal/router"
	"log"
	"net/http"
)

func main() {
	r := router.StRout()
	log.Println("Starting gateway on :8080...")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Failed to start gateway: %v", err)
	}
}
