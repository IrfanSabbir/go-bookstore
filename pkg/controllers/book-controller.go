package controllers

import (
	"fmt"

	models "github.com/IrfanSabbir/go-bookstore/pkg/models"
	utils "github.com/IrfanSabbir/go-bookstore/pkg/utils"

	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func getUserIdFromContext(r *http.Request) (int64, error) {
	ctx := r.Context()
	idString, ok := ctx.Value("user_id").(string)
	if !ok {
		return -1, fmt.Errorf("Unexpected id")
	}
	user_id, err := strconv.ParseInt(idString, 0, 0)
	if err != nil {
		return -1, fmt.Errorf("Unexpected id")
	}
	return user_id, nil

}

func GetBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("I am here")
	user_id, err := getUserIdFromContext(r)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	fmt.Println("user id", user_id)

	books := models.GetAllBooks()
	res, _ := json.Marshal(books)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}

	book := models.GetBookById(id)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	createBook := models.Book{}
	utils.ParseBody(r, &createBook)
	cratedBook := createBook.CreateBook()
	res, _ := json.Marshal(cratedBook)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	id, _ := strconv.ParseInt(bookId, 0, 0)
	deletedBook := models.DeleteBook(id)
	res, _ := json.Marshal(deletedBook)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	bookId := vars["bookId"]
	id, _ := strconv.ParseInt(bookId, 0, 0)
	updateBookItem := &models.Book{}
	utils.ParseBody(r, updateBookItem)
	updatedItem := updateBookItem.UpadteBook(id)
	res, _ := json.Marshal(updatedItem)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
