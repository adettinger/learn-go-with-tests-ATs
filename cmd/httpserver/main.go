package main

import (
	"log"
	"net/http"

	go_specs_greet "github.com/adettinger/learn-go-with-tests-ATs"
)

func main() {
	log.Println("Starting HTTP server...")
	handler := http.HandlerFunc(go_specs_greet.Handler)
	log.Println("Server started on :8080")
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal(err)
	}
}
