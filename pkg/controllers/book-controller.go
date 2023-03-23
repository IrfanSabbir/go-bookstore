package controllers

import (
	models "github.com/IrfanSabbir/go-bookstore/pkg/models"

	"encoding/json"
	"net/http"
)

func GetBook(w http.ResponseWriter, r *http.Request) {
	books := models.GetAllBooks()
	res, _ := json.Marshal(books)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
