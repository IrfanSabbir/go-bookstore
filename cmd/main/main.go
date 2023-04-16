package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/IrfanSabbir/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Could not load env")
	} else {
		fmt.Println("Loaded all env")
	}
}

func main() {
	r := mux.NewRouter()

	routes.RegisterBookRoutes(r)
	routes.RegisterUserRoutes(r)
	http.Handle("/", r)

	fmt.Printf("Starting server on 8080\n")
	log.Fatal(http.ListenAndServe(":8080", r))
}
