package routes

import (
	controllers "github.com/IrfanSabbir/go-bookstore/pkg/controllers"
	"github.com/gorilla/mux"
)

func RegisterUserRoutes(router *mux.Router) {
	router.HandleFunc("/register", controllers.RegisterUser).Methods("POST")
	router.HandleFunc("/login", controllers.Login).Methods("POST")
}
