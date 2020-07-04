package main

import (
	"log"
	"net/http"

	"github.com/nirbhaybagmar9/aws_go/handlers"
)

func main() {
	// this creates a new http.ServeMux, which is used to register handlers to execute in response to routes
	mux := http.NewServeMux()
	mux.Handle("/", handlers.GetInstances())
	mux.Handle("/state", handlers.ChangeState())
	mux.Handle("/create", handlers.CreateInstance())


	log.Printf("serving on port 8080")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
