package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/IrfanSabbir/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
)

type Message struct {
	Body string `json:"body"`
}

var message Message

func main() {
	r := mux.NewRouter()

	routes.RegisterBookRoutes(r)
	http.Handle("/", r)

	fmt.Printf("Starting server on 8080\n")
	log.Fatal(http.ListenAndServe(":8080", r))
}
