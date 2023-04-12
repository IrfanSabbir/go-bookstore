package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	jwt_token "github.com/IrfanSabbir/go-bookstore/pkg/jwt"
	model "github.com/IrfanSabbir/go-bookstore/pkg/models"
	"github.com/IrfanSabbir/go-bookstore/pkg/utils"
)

type AuthBody struct {
	email    string
	password string
}
type AuthUserResponse struct {
	User    map[string]interface{} `json:"user"`
	Message string                 `json:"message"`
	Token   string                 `json:"token"`
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
	fmt.Println("authUser", authUser)
	tokenString, err := jwt_token.GenerateToken(int64(authUser.ID))
	fmt.Println("tokenString", tokenString)

	if err != nil {
		fmt.Println(err)
	}
	respObject := AuthUserResponse{
		User: map[string]interface{}{
			"id":    authUser.ID,
			"name":  authUser.Name,
			"email": authUser.Email,
			"role":  authUser.Role,
		},
		Message: "Successfullt login",
		Token:   tokenString,
	}

	fmt.Println(respObject)
	res, err := json.Marshal(respObject)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
