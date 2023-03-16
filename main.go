package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Message struct {
	Body string `json:"body"`
}

var message Message

func main() {
	request := mux.NewRouter()

	request.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("INto this route")
		w.Header().Set("Content-Type", "application/json")

		message.Body = "I am in home"
		json.NewEncoder(w).Encode(message)

	}).Methods("GET")

	fmt.Printf("Starting server on 8080\n")
	log.Fatal(http.ListenAndServe(":8080", request))
}
