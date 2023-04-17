package routes

import (
	controllers "github.com/IrfanSabbir/go-bookstore/pkg/controllers"
	middleware "github.com/IrfanSabbir/go-bookstore/pkg/middleware"
	"github.com/gorilla/mux"
)

var RegisterBookRoutes = func(router *mux.Router) {
	router.HandleFunc("/books", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book", middleware.AuthMiddleware(controllers.CreateBook)).Methods("POST")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", middleware.AuthMiddleware(controllers.UpdateBook)).Methods("PUT")
	router.HandleFunc("/book/{bookId}", middleware.AuthMiddleware(controllers.DeleteBook)).Methods("DELETE")
}
