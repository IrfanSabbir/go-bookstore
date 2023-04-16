package middleware

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintln(w, "Missing authorization token")
			return
		}
		beareToken := strings.Split(tokenString, " ")[1]
		token, err := jwt.Parse(beareToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Invalid authorization token %v", token.Header["alg"])
			}
			jwt_secret := []byte(os.Getenv("API_SECRET"))
			return jwt_secret, nil
		})
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintln(w, err)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

			user_id := fmt.Sprint(claims["user_id"])

			ctx := context.WithValue(r.Context(), "user_id", user_id)
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			fmt.Println("FAcing errror here")

			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintln(w, "Invalid authorization token")
			return
		}
	})
}
