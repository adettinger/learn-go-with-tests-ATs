package main

import (
	"log"
	"net/http"

	"github.com/adettinger/learn-go-with-tests-ATs/adapters/httpserver"
)

func main() {
	log.Println("Starting HTTP server...")
	handler := httpserver.NewHandler()
	log.Println("Server started on :8080")
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal(err)
	}
}
