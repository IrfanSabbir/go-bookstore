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
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(err.Error()))
		return
	}

	book := models.GetBookById(id)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	user_id, err := getUserIdFromContext(r)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(err.Error()))
		return
	}

	createBook := models.Book{}
	utils.ParseBody(r, &createBook)
	cratedBook := createBook.CreateBook(user_id)
	res, _ := json.Marshal(cratedBook)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	user_id, err := getUserIdFromContext(r)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(err.Error()))
		return
	}

	vars := mux.Vars(r)
	bookId := vars["bookId"]
	id, _ := strconv.ParseInt(bookId, 0, 0)
	deletedBook, err := models.DeleteBook(id, user_id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte(err.Error()))
		return
	}
	res, _ := json.Marshal(deletedBook)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	user_id, err := getUserIdFromContext(r)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(err.Error()))
		return
	}
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	id, _ := strconv.ParseInt(bookId, 0, 0)
	updateBookItem := &models.Book{}
	utils.ParseBody(r, updateBookItem)
	updatedItem, err := updateBookItem.UpadteBook(id, user_id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte(err.Error()))
		return
	}
	res, _ := json.Marshal(updatedItem)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
