package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	model "github.com/IrfanSabbir/go-bookstore/pkg/models"
	"github.com/IrfanSabbir/go-bookstore/pkg/utils"
)

type AuthBody struct {
	email    string
	password string
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	registerUser := model.User{}
	utils.ParseBody(r, &registerUser)
	fmt.Println(registerUser)
	registeredUser := registerUser.RegisterUser()
	res, _ := json.Marshal(registeredUser)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func Login(w http.ResponseWriter, r *http.Request) {
	authBody := model.User{}
	utils.ParseBody(r, &authBody)
	authUser := model.Login(authBody.Email, authBody.Password)
	fmt.Println(authUser)
	res, _ := json.Marshal(authUser)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
